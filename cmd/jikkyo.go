package cmd

import (
	"encoding/xml"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/hogecode/JikkyoUtil/internal/api"
	"github.com/hogecode/JikkyoUtil/internal/presentation"
	"github.com/hogecode/JikkyoUtil/internal/usecase"
)

// runjikkyo is the main execution logic
func runjikkyo() error {
	// Validate inputs
	if title == "" {
		return errorf("title is required (use -t or --title)")
	}
	if episode <= 0 {
		return errorf("episode must be a positive number (use -e or --episode)")
	}

	// Setup logger
	loggerCfg := presentation.LoggerConfig{
		Verbose: verbose,
		Output:  os.Stderr,
		LogFile: logFile,
	}
	logger, err := presentation.NewLogger(loggerCfg)
	if err != nil {
		return errorf("failed to create logger: %v", err)
	}

	// Create API client
	client := api.NewClient()

	// Create core use case
	coreUC := usecase.NewCoreUseCase(client, logger, os.Stdin)

	// Execute workflow
	result, err := coreUC.Execute(title, episode)
	logger.Info("jikkyo workflow completed", slog.Any("result", result))
	if err != nil {
		return err
	}

	// Output result
	outputter := presentation.NewOutputFormatter(verbose)
	outputter.PrintResult(result)

	// Write program info file if output directory is specified
	if outputDir != "" && result.ProgramFileName != "" && result.ProgramContent != "" {
		filePath := filepath.Join(outputDir, result.ProgramFileName)
		// Convert to CRLF for Windows
		contentWithCRLF := []byte(strings.ReplaceAll(result.ProgramContent, "\r\n", "\n"))
		contentWithCRLF = []byte(strings.ReplaceAll(string(contentWithCRLF), "\n", "\r\n"))
		err := os.WriteFile(filePath, contentWithCRLF, 0644)
		if err != nil {
			logger.Error("failed to write program info file",
				slog.String("path", filePath),
				slog.String("error", err.Error()))
		} else {
			logger.Info("program info file written successfully",
				slog.String("path", filePath))
		}
	}

	logger.Info("starting jikkyo log file generation")
	// Write Jikkyo log file (XML format) if output directory is specified and we have the necessary info
	if outputDir != "" && result.ProgramFileName != "" && result.JikkyoID != ""  {
		logger.Info("fetching jikkyo logs from API",
			slog.String("jikkyo_id", result.JikkyoID),
			slog.Int64("start_time", result.StartTimeUnix),
			slog.Int64("end_time", result.EndTimeUnix))

		// Call Jikkyo API to fetch comments in XML format
		xmlResponse, err := client.GetJikkyoCommentsXML(result.JikkyoID, result.StartTimeUnix, result.EndTimeUnix)

		if err != nil {
			logger.Error("failed to fetch jikkyo comments from API",
				slog.String("error", err.Error()))
		} else {
			// Marshal the XML response back to bytes
			xmlContent, err := xml.MarshalIndent(xmlResponse, "", "  ")
			if err != nil {
				logger.Error("failed to marshal jikkyo XML response",
					slog.String("error", err.Error()))
			} else {
			logger.Debug("successfully fetched jikkyo comments in XML format",
				slog.Int("content_length", len(xmlContent)),
				slog.Int("packet_count", len(xmlResponse.Chats)))

				// Generate jikkyo filename from program filename
				jikkyoFileName := result.ProgramFileName
				if len(result.ProgramFileName) > len(".ts.program.txt") {
					// Replace .ts.program.txt with .xml
					jikkyoFileName = result.ProgramFileName[:len(result.ProgramFileName)-len(".ts.program.txt")] + ".xml"
				} else {
					// Fallback: just append .xml
					jikkyoFileName = result.ProgramFileName + ".xml"
				}

				jikkyoFilePath := filepath.Join(outputDir, jikkyoFileName)

				// Write XML content to file with CRLF for Windows
				xmlContentStr := string(xmlContent)
				xmlContentStr = strings.ReplaceAll(xmlContentStr, "\r\n", "\n")
				xmlContentStr = strings.ReplaceAll(xmlContentStr, "\n", "\r\n")
				xmlContentWithCRLF := []byte(xmlContentStr)
				
				err := os.WriteFile(jikkyoFilePath, xmlContentWithCRLF, 0644)
				if err != nil {
					logger.Error("failed to write jikkyo log file",
						slog.String("path", jikkyoFilePath),
						slog.String("error", err.Error()))
				} else {
					logger.Info("jikkyo log file written successfully",
						slog.String("path", jikkyoFilePath))
				}
			}
		}
	}

	return nil
}

// errorf returns an error with a formatted message
func errorf(msg string, args ...interface{}) error {
	return returnError{fmt.Sprintf(msg, args...)}
}

type returnError struct {
	msg string
}

func (e returnError) Error() string {
	return e.msg
}

package log

/*
	Sliver Implant Framework
	Copyright (C) 2022  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"io"

	"github.com/sirupsen/logrus"
)

// StartLogger - Returns a logger backed by the given io.Writer
func StartLogger(backingFile io.Writer) *logrus.Logger {
	txtLogger := logrus.New()
	txtLogger.Formatter = &logrus.TextFormatter{
		ForceColors:   false,
		FullTimestamp: true,
	}
	txtLogger.Out = backingFile
	txtLogger.SetLevel(logrus.DebugLevel)
	return txtLogger
}

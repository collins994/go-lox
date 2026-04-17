@echo off

set "ACTION=%1"
set "EXECUTABLE_FILE=go_lox.exe"
set "GIT_BRANCH="

if "%ACTION%" == ""  (
	echo USAGE: make build ^| run ^[file^]
	EXIT /B
)

if %ACTION% EQU build (
	echo building into .\build\...
	go build -o .\build\%EXECUTABLE_FILE% .\code && echo done! || echo compilation failed, error code: %ERRORLEVEL%
	EXIT /B
) 

if %ACTION% EQU test (
	echo  testing .\code...
	go test .\code
	EXIT /B
) 

if %ACTION% EQU run (
	.\build\%EXECUTABLE_FILE% %2 
	EXIT /B
)

if %ACTION% EQU ast (
	go run tool\generateAST.go ast
	EXIT /B
)

if %ACTION% EQU tags (
	echo generating tags...
	ctags -R  --exclude=build
	EXIT /B
)

if %ACTION% EQU push (
	git push https://github.com/collins994/go-lox.git %GIT_BRANCH%
	EXIT /B
)

if %ACTION% EQU pull (
	git pull https://github.com/collins994/go-lox.git
	EXIT /B
)

if %ACTION% EQU clean (
	echo cleaning the build folder...
	del .\build\*
	echo done!
	EXIT /B
)


echo invalid option "%ACTION%"

@echo off
wails build -clean -debug
copy .env build\bin\.env
SET GOPATH=%CD%\gopath
SET PATH=C:\Go\bin;C:\Program Files\Git\cmd\;%GOPATH%\bin;%PATH%

cd %GOPATH%\src\github.com\cloudfoundry\cli

powershell -command set-executionpolicy remotesigned

go get github.com/onsi/ginkgo/ginkgo

ginkgo -r -randomizeAllSpecs .

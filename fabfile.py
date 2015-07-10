from fabric.api import *
import cuisine
import os

project_path = "src/bi"

def deploy():
	path = run("echo $GOPATH")
	path = os.path.join(path, "src", "gitlab.com", "warnabroda", "warnabrodagomartini")
	with cd(path):
		run("git pull")
		run("go get")
		run("go build")
		run("mv warnabrodagomartini /opt/warnabroda/project/server/server")
		run("cp -f tools/* /opt/warnabroda/project/server/tools")
		run("cp -f resource/* /opt/warnabroda/project/server/resource")
		run("cp -f messages/*yaml /opt/warnabroda/project/server/messages")

def deploy_tools():
	path = run("echo $GOPATH")
	path = os.path.join(path, "src", "gitlab.com", "warnabroda", "warnabrodagomartini")
	with cd(path):
		run("git pull")
		run("cp tools/* /opt/warnabroda/project/server/tools")

	run("supervisorctl restart whats_python")

def restart():
	run("supervisorctl restart server")


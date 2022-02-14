package executor

import "stvcv2/commandFormatter"

type ExecService struct {
	sequenceId    int
	applicationId int
	deviceId      int
	headCommand   *commandFormatter.Command
	result        []commandFormatter.CommandResult
}

func (E *ExecService) Init(sequenceId int, applicationId int, deviceId int) {
	E.sequenceId = sequenceId
	E.applicationId = applicationId
	E.deviceId = deviceId
}

func (E *ExecService) Run() { E.prepare().exec().setResult() }

func (E *ExecService) GetResult() []commandFormatter.CommandResult { return E.result }

func (E *ExecService) setResult() *ExecService {
	E.headCommand.GetResult(&E.result)
	return E
}
func (E *ExecService) prepare() *ExecService {
	E.headCommand = commandFormatter.FormatQuery()
	return E
}

func (E *ExecService) exec() *ExecService {
	E.headCommand.Exec()
	return E
}

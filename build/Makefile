BINARY_NAME=../bin
KUBE_CONTROLLER_MANAGER=../cmd/kube-controller-manager/controller-manager.go
API_SERVER=../cmd/kube-apiserver/apiserver.go
SCHEDULER=../cmd/kube-scheduler/scheduler.go
KUBELET=../cmd/kubelet/kubeletMain.go

build_windows:
	go build -o ${BINARY_NAME}/controller.exe ${KUBE_CONTROLLER_MANAGER}
	go build -o ${BINARY_NAME}/apiserver.exe ${API_SERVER}
	go build -o ${BINARY_NAME}/scheduler.exe ${SCHEDULER}
	go build -o ${BINARY_NAME}/kubelet.exe ${KUBELET}


build_linux:
	rm ${BINARY_NAME}/*
	go build -o ${BINARY_NAME}/controller ${KUBE_CONTROLLER_MANAGER}
	go build -o ${BINARY_NAME}/apiserver ${API_SERVER}
	go build -o ${BINARY_NAME}/scheduler ${SCHEDULER}
	go build -o ${BINARY_NAME}/kubelet ${KUBELET}








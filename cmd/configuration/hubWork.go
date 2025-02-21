package configuration

import (
	"crave/hub/cmd/work/cmd/api/domain/service"
	"crave/hub/cmd/work/cmd/api/infrastructure/cache"
	"crave/hub/cmd/work/cmd/api/presentation/controller"
	"net/http"
)

type HubWorkContainer struct {
	Variable       *Variable
	WorkController controller.IController
	WorkService    service.IService
	WorkCache      cache.ICache
}

// defineDatabase implements configuration.IContainer.
func (ctnr *HubWorkContainer) DefineDatabase() error {
	return nil
}

// getHttpHandler implements configuration.IContainer.
func (ctnr *HubWorkContainer) GetHttpHandler() http.Handler {
	return nil
}

// initVariable implements configuration.IContainer.
func (ctnr *HubWorkContainer) InitVariable() {
	return
}

// setRouter implements configuration.IContainer.
func (ctnr *HubWorkContainer) SetRouter(router any) {
	return
}

// DefineGrpc implements configuration.IContainer.
func (ctnr *HubWorkContainer) DefineGrpc() error {
	// lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ctnr.Variable.ctnrubApiIp, ctnr.Variable.ctnrubApiPort))
	// if err != nil {
	// 	return fmt.Errorf("failed to listen : %d, %w", ctnr.Variable.ctnrubApiPort, err)
	// }
	// s := grpc.NewServer()

	// if servErr := s.Serve(lis); servErr != nil {
	// 	return fmt.Errorf("failed to create server: %w", err)
	// }
	return nil
}

func (ctnr *HubWorkContainer) DefineRoute() error {
	return nil
}

func (ctnr *HubWorkContainer) InitDependency(dependency any) {
	ctnr.WorkCache = cache.NewCache()
	ctnr.WorkService = service.NewService(ctnr.WorkCache)
}

func NewHubWorkContainer() *HubWorkContainer {
	ctnr := &HubWorkContainer{}
	ctnr.InitVariable()
	ctnr.InitDependency(nil)
	return ctnr
}

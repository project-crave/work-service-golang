package configuration

import (
	"crave/hub/cmd/model"
	"crave/hub/cmd/work/cmd/api/domain/service"
	"crave/hub/cmd/work/cmd/api/infrastructure/cache"
	"crave/hub/cmd/work/cmd/api/infrastructure/repository"
	"crave/hub/cmd/work/cmd/api/presentation/controller"
	"crave/shared/database"
	"net/http"
)

type HubWorkContainer struct {
	Variable       *Variable
	MysqlWrapper   *database.MysqlWrapper
	WorkController controller.IController
	WorkService    service.IService
	WorkCache      cache.ICache
	WorkRepository repository.IRepository
}

// defineDatabase implements configuration.IContainer.
func (ctnr *HubWorkContainer) DefineDatabase() error {
	ctnr.MysqlWrapper = database.ConnectMysqlDatabase(ctnr.Variable.Database)
	if err := ctnr.MysqlWrapper.Driver.AutoMigrate(&model.Work{}); err != nil {
		return err
	}
	return nil
}

// getHttpHandler implements configuration.IContainer.
func (ctnr *HubWorkContainer) GetHttpHandler() http.Handler {
	return nil
}

// initVariable implements configuration.IContainer.
func (ctnr *HubWorkContainer) InitVariable() error {
	ctnr.Variable = NewVariable()
	return nil
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

func (ctnr *HubWorkContainer) InitDependency(dependency any) error {
	ctnr.DefineDatabase()
	ctnr.WorkRepository = repository.NewRepository(ctnr.MysqlWrapper)
	ctnr.WorkCache = cache.NewCache()
	ctnr.WorkService = service.NewService(ctnr.WorkCache, ctnr.WorkRepository)
	return nil
}

func NewHubWorkContainer() *HubWorkContainer {
	ctnr := &HubWorkContainer{}
	ctnr.InitVariable()
	ctnr.InitDependency(nil)
	return ctnr
}

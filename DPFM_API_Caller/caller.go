package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-article-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-article-creates-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-article-creates-rmq-kube/config"
	"data-platform-api-article-creates-rmq-kube/sub_func_complementer"
	database "github.com/latonaio/golang-mysql-network-connector"
	"sync"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type DPFMAPICaller struct {
	ctx  context.Context
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient

	//configure    *existence_conf.ExistenceConf
	complementer *sub_func_complementer.SubFuncComplementer

	db *database.Mysql
}

func NewDPFMAPICaller(
	conf *config.Conf, rmq *rabbitmq.RabbitmqClient,
	//confirmor *existence_conf.ExistenceConf,
	complementer *sub_func_complementer.SubFuncComplementer,
	db *database.Mysql,
) *DPFMAPICaller {
	return &DPFMAPICaller{
		ctx:  context.Background(),
		conf: conf,
		rmq:  rmq,
		//configure:    confirmor,
		complementer: complementer,
		db:           db,
	}
}

func (c *DPFMAPICaller) AsyncCreates(
	accepter []string,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	log *logger.Logger,
) (interface{}, []error) {
	//wg := sync.WaitGroup{}
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)
	//exconfAllExist := false

	//exconfFin := make(chan error)
	//subFuncFin := make(chan error)

	subfuncSDC := &sub_func_complementer.SDC{}

	// 他PODへ問い合わせ
	//wg.Add(1)
	//go c.exconfProcess(&mtx, &wg, exconfFin, input, output, &exconfAllExist, accepter, &errs, log)
	//if input.APIType == "creates" {
	//	go c.subfuncProcess(&mtx, &wg, subFuncFin, input, output, subfuncSDC, accepter, &errs, log)
	//} else if input.APIType == "updates" {
	//	go func() { subFuncFin <- nil }()
	//} else {
	//	go func() { subFuncFin <- nil }()
	//}

	// 処理待ち
	//ticker := time.NewTicker(10 * time.Second)
	//if err := c.finWait(&mtx, exconfFin, ticker); err != nil || len(errs) != 0 {
	//	if err != nil {
	//		errs = append(errs, err)
	//	}
	//	return nil, errs
	//}
	//if !exconfAllExist {
	//	mtx.Lock()
	//	return nil, nil
	//}
	//wg.Wait()
	//if input.APIType == "creates" {
	//	for range accepter {
	//		if err := c.finWait(&mtx, subFuncFin, ticker); err != nil || len(errs) != 0 {
	//			if err != nil {
	//				errs = append(errs, err)
	//			}
	//			return subfuncSDC.Message, errs
	//		}
	//	}
	//} else if input.APIType == "updates" {
	//	if err := c.finWait(&mtx, subFuncFin, ticker); err != nil || len(errs) != 0 {
	//		if err != nil {
	//			errs = append(errs, err)
	//		}
	//		return subfuncSDC.Message, errs
	//	}
	//}

	var response interface{}
	// SQL処理
	if input.APIType == "creates" {
		response = c.createSqlProcess(nil, &mtx, input, output, subfuncSDC, accepter, &errs, log)
	} else if input.APIType == "updates" {
		response = c.updateSqlProcess(nil, &mtx, input, output, accepter, &errs, log)
	}

	return response, nil
}

func (c *DPFMAPICaller) exconfProcess(
	mtx *sync.Mutex,
	wg *sync.WaitGroup,
	exconfFin chan error,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	exconfAllExist *bool,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) {
	defer wg.Done()
	var e []error
	//*exconfAllExist, e = c.configure.Conf(input, output, accepter, log)
	if len(e) != 0 {
		mtx.Lock()
		*errs = append(*errs, e...)
		mtx.Unlock()
		exconfFin <- xerrors.New("exconf error")
		return
	}
	exconfFin <- nil
}

func (c *DPFMAPICaller) subfuncProcess(
	mtx *sync.Mutex,
	wg *sync.WaitGroup,
	subFuncFin chan error,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) {
	for _, fn := range accepter {
		wg.Add(1)
		switch fn {
		case "Header":
			c.headerCreate(mtx, wg, subFuncFin, input, output, subfuncSDC, errs, log)
		case "Partner":
			c.partnerCreate(mtx, wg, subFuncFin, input, output, subfuncSDC, errs, log)
		case "Address":
			c.addressCreate(mtx, wg, subFuncFin, input, output, subfuncSDC, errs, log)
		case "Counter":
			c.counterCreate(mtx, wg, subFuncFin, input, output, subfuncSDC, errs, log)
		case "Like":
			c.likeCreate(mtx, wg, subFuncFin, input, output, subfuncSDC, errs, log)
		default:
			wg.Done()
		}
	}
}

func (c *DPFMAPICaller) finWait(
	mtx *sync.Mutex,
	finChan chan error,
	ticker *time.Ticker,
) error {
	select {
	case e := <-finChan:
		if e != nil {
			mtx.Lock()
			return e
		}
	case <-ticker.C:
		return xerrors.New("time out")
	}
	return nil
}

func (c *DPFMAPICaller) headerCreate(
	mtx *sync.Mutex,
	wg *sync.WaitGroup,
	errFin chan error,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) {
	var err error = nil
	defer func() {
		errFin <- err
	}()
	defer wg.Done()
	err = c.complementer.ComplementHeader(input, subfuncSDC, log)
	if err != nil {
		mtx.Lock()
		*errs = append(*errs, err)
		mtx.Unlock()
		return
	}
	output.SubfuncResult = getBoolPtr(true)
	if subfuncSDC.SubfuncResult == nil || !*subfuncSDC.SubfuncResult {
		output.SubfuncResult = getBoolPtr(false)
		output.SubfuncError = subfuncSDC.SubfuncError
		err = xerrors.New(output.SubfuncError)
		return
	}
	return
}

func (c *DPFMAPICaller) partnerCreate(
	mtx *sync.Mutex,
	wg *sync.WaitGroup,
	errFin chan error,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) {
	//var err error = nil
	//defer func() {
	//	errFin <- err
	//}()
	//defer wg.Done()
	//err = c.complementer.ComplementPartner(input, subfuncSDC, log)
	//if err != nil {
	//	mtx.Lock()
	//	*errs = append(*errs, err)
	//	mtx.Unlock()
	//	return
	//}
	//output.SubfuncResult = getBoolPtr(true)
	//if subfuncSDC.SubfuncResult == nil || !*subfuncSDC.SubfuncResult {
	//	output.SubfuncResult = getBoolPtr(false)
	//	output.SubfuncError = subfuncSDC.SubfuncError
	//	err = xerrors.New(output.SubfuncError)
	//	return
	//}

	return
}

func (c *DPFMAPICaller) addressCreate(
	mtx *sync.Mutex,
	wg *sync.WaitGroup,
	errFin chan error,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) {
	//var err error = nil
	//defer func() {
	//	errFin <- err
	//}()
	//defer wg.Done()
	//err = c.complementer.ComplementAddress(input, subfuncSDC, log)
	//if err != nil {
	//	mtx.Lock()
	//	*errs = append(*errs, err)
	//	mtx.Unlock()
	//	return
	//}
	//output.SubfuncResult = getBoolPtr(true)
	//if subfuncSDC.SubfuncResult == nil || !*subfuncSDC.SubfuncResult {
	//	output.SubfuncResult = getBoolPtr(false)
	//	output.SubfuncError = subfuncSDC.SubfuncError
	//	err = xerrors.New(output.SubfuncError)
	//	return
	//}

	return
}

func (c *DPFMAPICaller) counterCreate(
	mtx *sync.Mutex,
	wg *sync.WaitGroup,
	errFin chan error,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) {
	//var err error = nil
	//defer func() {
	//	errFin <- err
	//}()
	//defer wg.Done()
	//err = c.complementer.ComplementCounter(input, subfuncSDC, log)
	//if err != nil {
	//	mtx.Lock()
	//	*errs = append(*errs, err)
	//	mtx.Unlock()
	//	return
	//}
	//output.SubfuncResult = getBoolPtr(true)
	//if subfuncSDC.SubfuncResult == nil || !*subfuncSDC.SubfuncResult {
	//	output.SubfuncResult = getBoolPtr(false)
	//	output.SubfuncError = subfuncSDC.SubfuncError
	//	err = xerrors.New(output.SubfuncError)
	//	return
	//}

	return
}

func (c *DPFMAPICaller) likeCreate(
	mtx *sync.Mutex,
	wg *sync.WaitGroup,
	errFin chan error,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) {
	//var err error = nil
	//defer func() {
	//	errFin <- err
	//}()
	//defer wg.Done()
	//err = c.complementer.ComplementLike(input, subfuncSDC, log)
	//if err != nil {
	//	mtx.Lock()
	//	*errs = append(*errs, err)
	//	mtx.Unlock()
	//	return
	//}
	//output.SubfuncResult = getBoolPtr(true)
	//if subfuncSDC.SubfuncResult == nil || !*subfuncSDC.SubfuncResult {
	//	output.SubfuncResult = getBoolPtr(false)
	//	output.SubfuncError = subfuncSDC.SubfuncError
	//	err = xerrors.New(output.SubfuncError)
	//	return
	//}

	return
}

func checkResult(msg rabbitmq.RabbitmqMessage) bool {
	data := msg.Data()

	sqlCreatesResult := data["result"]

	if sqlCreatesResult == "success" {
		return true
	} else {
		return false
	}
}

func getBoolPtr(b bool) *bool {
	return &b
}

func contains(slice []string, target string) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

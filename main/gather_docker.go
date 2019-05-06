package main

import (
	"math"
	"metric_monitor/common/register"
	"metric_monitor/common/types"
	"metric_monitor/itom_agent/g"
	"metric_monitor/itom_agent/mgr"
	"metric_monitor/itom_agent/models/metric"
)

func gatherDocker(uuid string, conf *register.AgentBasicConf, jobid string) (am []*types.AgentMetric) {
	if conf == nil {
		return
	}
	// 在实时数据没有采集时上报docker静态指标供巡检平台查询
	if jobid != "" && mgr.ConfigManager.GetAgtMetricsConf().DockerConf == nil {
		containers, _ := metric.GetContainerInfo()
		dockerMutex.Lock()
		dockerGlob = make([]register.ContainerInfo, 0)
		for _, ctn := range containers {
			ctni := register.ContainerInfo{}
			ctni.ContainerId = ctn.ContainerId
			ctni.ContainerName = ctn.ContainerName
			ctni.ContainerLabels = ctn.ContainerLabels
			ctni.ContainerStatus = ctn.ContainerStatus
			ctni.ContainerEnv = ctn.ContainerEnv
			ctni.ContainerIP = ctn.ContainerIP
			dockerGlob = append(dockerGlob, ctni)
		}
		dockerMutex.Unlock()
		go registerAgent(uuid)
	}

	timestamp := time.Now()
	// 获取容器实时指标
	ctn, err := metric.GetContainerStats()
	if err != nil {
		log.Error("get docker stats: %v", err)
		if jobid == "" {
			return
		}
		if metric.IsTimeout(err) {
			ctn[""] = &types.ContainerStats{ErrType: errTimeout}
		} else {
			ctn[""] = &types.ContainerStats{ErrType: errFail}
		}
	}

	for _, m := range conf.Metrics {
		for k, v := range ctn {
			if v == nil {
				continue
			}
			tags := make(map[string]string)
			if jobid != "" {
				tags[jobIdKey] = jobid
			}
			tags[g.MTKeyName] = k
			md5sum := tagsMd5sum(tags)
			tags["md5"] = md5sum
			metric := &types.AgentMetric{
				Uuid:   uuid,
				Metric: m,
				Tags:   tags,
				//Endpoint: g.Hostname,
				Clock: timestamp.Unix(),
				Date:  timestamp}
			if v.ErrType == errTimeout {
				metric.Value = math.MinInt32
			} else if v.ErrType == errFail {
				metric.Value = nil
			} else {
				switch m {
				case types.MetricName[types.CCpuUsage]:
					metric.Value = v.ContainerCpuUsage
				case types.MetricName[types.CMemUsage]:
					metric.Value = v.ContainerMemUsage
				case types.MetricName[types.CTotalByteInSec]:
					metric.Value = v.ContainerNetBytesInSec
				case types.MetricName[types.CTotalByteOutSec]:
					metric.Value = v.ContainerNetBytesOutSec
				}
			}
			am = append(am, metric)
			log.Debug("docker metric :%+v", metric)
		}
	}

	return am
}

func main() {

}

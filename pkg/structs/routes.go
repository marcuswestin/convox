package structs

var routes = map[string]string{}

func init() {
	routes["Initialize"] = ""
	routes["Start"] = ""
	routes["AppCancel"] = "POST /apps/{name}/cancel"
	routes["AppCreate"] = "POST /apps"
	routes["AppDelete"] = "DELETE /apps/{name}"
	routes["AppGet"] = "GET /apps/{name}"
	routes["AppList"] = "GET /apps"
	routes["AppLogs"] = "SOCKET /apps/{name}/logs"
	routes["AppMetrics"] = "GET /apps/{name}/metrics"
	routes["AppUpdate"] = "PUT /apps/{name}"
	routes["BalancerList"] = "GET /apps/{app}/balancers"
	routes["BuildCreate"] = "POST /apps/{app}/builds"
	routes["BuildExport"] = "GET /apps/{app}/builds/{id}.tgz"
	routes["BuildGet"] = "GET /apps/{app}/builds/{id}"
	routes["BuildImport"] = "POST /apps/{app}/builds/import"
	routes["BuildLogs"] = "SOCKET /apps/{app}/builds/{id}/logs"
	routes["BuildList"] = "GET /apps/{app}/builds"
	routes["BuildUpdate"] = "PUT /apps/{app}/builds/{id}"
	routes["CapacityGet"] = "GET /system/capacity"
	routes["CertificateApply"] = "PUT /apps/{app}/ssl/{service}/{port}"
	routes["CertificateCreate"] = "POST /certificates"
	routes["CertificateDelete"] = "DELETE /certificates/{id}"
	routes["CertificateGenerate"] = "POST /certificates/generate"
	routes["CertificateList"] = "GET /certificates"
	routes["EventSend"] = "POST /events"
	routes["FilesDelete"] = "DELETE /apps/{app}/processes/{pid}/files"
	routes["FilesDownload"] = "GET /apps/{app}/processes/{pid}/files"
	routes["FilesUpload"] = "POST /apps/{app}/processes/{pid}/files"
	routes["InstanceKeyroll"] = "POST /instances/keyroll"
	routes["InstanceList"] = "GET /instances"
	routes["InstanceShell"] = "SOCKET /instances/{id}/shell"
	routes["InstanceTerminate"] = "DELETE /instances/{id}"
	routes["ObjectDelete"] = "DELETE /apps/{app}/objects/{key:.*}"
	routes["ObjectExists"] = "HEAD /apps/{app}/objects/{key:.*}"
	routes["ObjectFetch"] = "GET /apps/{app}/objects/{key:.*}"
	routes["ObjectList"] = "GET /apps/{app}/objects"
	routes["ObjectStore"] = "POST /apps/{app}/objects/{key:.*}"
	routes["ProcessExec"] = "SOCKET /apps/{app}/processes/{pid}/exec"
	routes["ProcessGet"] = "GET /apps/{app}/processes/{pid}"
	routes["ProcessList"] = "GET /apps/{app}/processes"
	routes["ProcessLogs"] = "SOCKET /apps/{app}/processes/{pid}/logs"
	routes["ProcessRun"] = "POST /apps/{app}/services/{service}/processes"
	routes["ProcessStop"] = "DELETE /apps/{app}/processes/{pid}"
	routes["Proxy"] = "SOCKET /proxy/{host}/{port}"
	routes["ReleaseCreate"] = "POST /apps/{app}/releases"
	routes["ReleaseGet"] = "GET /apps/{app}/releases/{id}"
	routes["ReleaseList"] = "GET /apps/{app}/releases"
	routes["ReleasePromote"] = "POST /apps/{app}/releases/{id}/promote"
	routes["RegistryAdd"] = "POST /registries"
	routes["RegistryList"] = "GET /registries"
	routes["RegistryRemove"] = "DELETE /registries/{server:.*}"
	routes["ResourceConsole"] = "SOCKET /apps/{app}/resources/{name}/console"
	routes["ResourceExport"] = "GET /apps/{app}/resources/{name}/data"
	routes["ResourceGet"] = "GET /apps/{app}/resources/{name}"
	routes["ResourceImport"] = "PUT /apps/{app}/resources/{name}/data"
	routes["ResourceList"] = "GET /apps/{app}/resources"
	routes["ServiceList"] = "GET /apps/{app}/services"
	routes["ServiceRestart"] = "POST /apps/{app}/services/{name}/restart"
	routes["ServiceUpdate"] = "PUT /apps/{app}/services/{name}"
	routes["SystemGet"] = "GET /system"
	routes["SystemLogs"] = "SOCKET /system/logs"
	routes["SystemInstall"] = ""
	routes["SystemMetrics"] = "GET /system/metrics"
	routes["SystemProcesses"] = "GET /system/processes"
	routes["SystemReleases"] = "GET /system/releases"
	routes["SystemResourceCreate"] = "POST /resources"
	routes["SystemResourceDelete"] = "DELETE /resources/{name}"
	routes["SystemResourceGet"] = "GET /resources/{name}"
	routes["SystemResourceLink"] = "POST /resources/{name}/links"
	routes["SystemResourceList"] = "GET /resources"
	routes["SystemResourceTypes"] = "OPTIONS /resources"
	routes["SystemResourceUnlink"] = "DELETE /resources/{name}/links/{app}"
	routes["SystemResourceUpdate"] = "PUT /resources/{name}"
	routes["SystemUninstall"] = ""
	routes["SystemUpdate"] = "PUT /system"
	routes["Workers"] = ""
}

func Routes() map[string]string {
	return routes
}

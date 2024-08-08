
/**
 * Samples for ContainerAppsDiagnostics GetDetector.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/ContainerAppsDiagnostics_Get.json
     */
    /**
     * Sample code: Get Container App's diagnostics info.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        getContainerAppSDiagnosticsInfo(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsDiagnostics().getDetectorWithResponse("mikono-workerapp-test-rg", "mikono-capp-stage1",
            "cappcontainerappnetworkIO", com.azure.core.util.Context.NONE);
    }
}

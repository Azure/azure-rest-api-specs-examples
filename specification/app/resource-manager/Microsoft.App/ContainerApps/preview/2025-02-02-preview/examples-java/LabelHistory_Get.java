
/**
 * Samples for ContainerAppsLabelHistory GetLabelHistory.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/
     * LabelHistory_Get.json
     */
    /**
     * Sample code: Get Container App's single label history.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        getContainerAppSSingleLabelHistory(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsLabelHistories().getLabelHistoryWithResponse("rg", "testContainerApp", "dev",
            com.azure.core.util.Context.NONE);
    }
}

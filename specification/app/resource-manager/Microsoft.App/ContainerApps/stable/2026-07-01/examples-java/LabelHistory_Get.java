
/**
 * Samples for ContainerAppsLabelHistory GetLabelHistory.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/LabelHistory_Get.json
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


/**
 * Samples for IntegrationFabrics Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/IntegrationFabrics_Get.json
     */
    /**
     * Sample code: IntegrationFabrics_Get.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void integrationFabricsGet(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.integrationFabrics().getWithResponse("myResourceGroup", "myWorkspace", "sampleIntegration",
            com.azure.core.util.Context.NONE);
    }
}

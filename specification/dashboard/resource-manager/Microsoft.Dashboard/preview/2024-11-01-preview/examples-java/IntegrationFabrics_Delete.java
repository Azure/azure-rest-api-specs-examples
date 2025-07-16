
/**
 * Samples for IntegrationFabrics Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/IntegrationFabrics_Delete.json
     */
    /**
     * Sample code: IntegrationFabrics_Delete.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void integrationFabricsDelete(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.integrationFabrics().delete("myResourceGroup", "myWorkspace", "sampleIntegration",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for IntegrationFabrics List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/IntegrationFabrics_List.json
     */
    /**
     * Sample code: IntegrationFabrics_List.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void integrationFabricsList(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.integrationFabrics().list("myResourceGroup", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}

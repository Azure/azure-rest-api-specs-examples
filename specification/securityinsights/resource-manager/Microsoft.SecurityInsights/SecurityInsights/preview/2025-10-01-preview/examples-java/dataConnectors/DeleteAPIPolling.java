
/**
 * Samples for DataConnectors Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/DeleteAPIPolling.json
     */
    /**
     * Sample code: Delete a APIPolling data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        deleteAAPIPollingDataConnector(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().deleteWithResponse("myRg", "myWorkspace", "316ec55e-7138-4d63-ab18-90c8a60fd1c8",
            com.azure.core.util.Context.NONE);
    }
}

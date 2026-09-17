
/**
 * Samples for DataConnectors List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/GetDataConnectors.json
     */
    /**
     * Sample code: Get all data connectors.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAllDataConnectors(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().list("myRg", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}

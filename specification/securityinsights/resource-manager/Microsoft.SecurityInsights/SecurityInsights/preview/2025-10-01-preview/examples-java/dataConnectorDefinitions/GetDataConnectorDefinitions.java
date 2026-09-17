
/**
 * Samples for DataConnectorDefinitions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectorDefinitions/GetDataConnectorDefinitions.json
     */
    /**
     * Sample code: Get all data connector definitions.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAllDataConnectorDefinitions(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectorDefinitions().list("myRg", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}

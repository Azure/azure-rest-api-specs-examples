
/**
 * Samples for DataConnectors Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/GetRestApiPollerById.json
     */
    /**
     * Sample code: Get a RestApiPoller data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getARestApiPollerDataConnector(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().getWithResponse("myRg", "myWorkspace",
            "RestApiPoller_fce27b90-d6f5-4d30-991a-af509a2b50a1", com.azure.core.util.Context.NONE);
    }
}

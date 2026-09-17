
/**
 * Samples for DataConnectors Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/GetGoogleCloudPlatformById.json
     */
    /**
     * Sample code: Get a GCP data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAGCPDataConnector(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().getWithResponse("myRg", "myWorkspace", "GCP_fce27b90-d6f5-4d30-991a-af509a2b50a1",
            com.azure.core.util.Context.NONE);
    }
}

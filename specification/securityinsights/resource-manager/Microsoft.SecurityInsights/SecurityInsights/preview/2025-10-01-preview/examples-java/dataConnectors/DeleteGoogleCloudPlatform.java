
/**
 * Samples for DataConnectors Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/DeleteGoogleCloudPlatform.json
     */
    /**
     * Sample code: Delete a GCP data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        deleteAGCPDataConnector(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().deleteWithResponse("myRg", "myWorkspace", "GCP_fce27b90-d6f5-4d30-991a-af509a2b50a1",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for DataConnectors Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/GetAzureAdvancedThreatProtectionById.json
     */
    /**
     * Sample code: Get an AATP data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAnAATPDataConnector(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().getWithResponse("myRg", "myWorkspace", "07e42cb3-e658-4e90-801c-efa0f29d3d44",
            com.azure.core.util.Context.NONE);
    }
}

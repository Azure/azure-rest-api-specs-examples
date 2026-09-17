
/**
 * Samples for DataConnectors Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/GetMicrosoftDefenderAdvancedThreatProtectionById.json
     */
    /**
     * Sample code: Get a MDATP data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAMDATPDataConnector(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().getWithResponse("myRg", "myWorkspace", "06b3ccb8-1384-4bcc-aec7-852f6d57161b",
            com.azure.core.util.Context.NONE);
    }
}

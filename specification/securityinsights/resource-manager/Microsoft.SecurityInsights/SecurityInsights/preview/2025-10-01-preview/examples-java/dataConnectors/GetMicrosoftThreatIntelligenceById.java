
/**
 * Samples for DataConnectors Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/GetMicrosoftThreatIntelligenceById.json
     */
    /**
     * Sample code: Get a MicrosoftThreatIntelligence data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAMicrosoftThreatIntelligenceDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().getWithResponse("myRg", "myWorkspace", "c345bf40-8509-4ed2-b947-50cb773aaf04",
            com.azure.core.util.Context.NONE);
    }
}

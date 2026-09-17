
/**
 * Samples for DataConnectors Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/DeleteMicrosoftThreatIntelligenceDataConnector.json
     */
    /**
     * Sample code: Delete an MicrosoftThreatIntelligence data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteAnMicrosoftThreatIntelligenceDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().deleteWithResponse("myRg", "myWorkspace", "c345bf40-8509-4ed2-b947-50cb773aaf04",
            com.azure.core.util.Context.NONE);
    }
}

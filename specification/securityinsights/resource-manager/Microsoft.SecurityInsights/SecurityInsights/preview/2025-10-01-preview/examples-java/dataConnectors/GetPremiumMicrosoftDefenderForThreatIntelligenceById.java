
/**
 * Samples for DataConnectors Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/GetPremiumMicrosoftDefenderForThreatIntelligenceById.json
     */
    /**
     * Sample code: Get a PremiumMicrosoftDefenderForThreatIntelligence data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAPremiumMicrosoftDefenderForThreatIntelligenceDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().getWithResponse("myRg", "myWorkspace", "8c569548-a86c-4fb4-8ae4-d1e35a6146f8",
            com.azure.core.util.Context.NONE);
    }
}

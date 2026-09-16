
/**
 * Samples for DataConnectors Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-10-01-preview/dataConnectors/DeletePremiumMicrosoftDefenderForThreatIntelligenceDataConnector.json
     */
    /**
     * Sample code: Deletes a PremiumMicrosoftDefenderForThreatIntelligence data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deletesAPremiumMicrosoftDefenderForThreatIntelligenceDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().deleteWithResponse("myRg", "myWorkspace", "8c569548-a86c-4fb4-8ae4-d1e35a6146f8",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for ThreatIntelligenceIndicatorsOperation List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/threatintelligence/GetThreatIntelligence.json
     */
    /**
     * Sample code: Get all threat intelligence indicators.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAllThreatIntelligenceIndicators(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.threatIntelligenceIndicatorsOperations().list("myRg", "myWorkspace", null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}

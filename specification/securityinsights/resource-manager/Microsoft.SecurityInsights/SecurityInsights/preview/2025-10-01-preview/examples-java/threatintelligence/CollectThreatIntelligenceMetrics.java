
/**
 * Samples for ThreatIntelligenceIndicatorMetrics List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/threatintelligence/CollectThreatIntelligenceMetrics.json
     */
    /**
     * Sample code: Get threat intelligence indicators metrics.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getThreatIntelligenceIndicatorsMetrics(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.threatIntelligenceIndicatorMetrics().listWithResponse("myRg", "myWorkspace",
            com.azure.core.util.Context.NONE);
    }
}

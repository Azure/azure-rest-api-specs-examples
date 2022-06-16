import com.azure.core.util.Context;

/** Samples for ThreatIntelligenceIndicatorMetrics List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/threatintelligence/CollectThreatIntelligenceMetrics.json
     */
    /**
     * Sample code: Get threat intelligence indicators metrics.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getThreatIntelligenceIndicatorsMetrics(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.threatIntelligenceIndicatorMetrics().listWithResponse("myRg", "myWorkspace", Context.NONE);
    }
}


/**
 * Samples for ThreatIntelligenceIndicator Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * threatintelligence/DeleteThreatIntelligence.json
     */
    /**
     * Sample code: Delete a threat intelligence indicator.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        deleteAThreatIntelligenceIndicator(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.threatIntelligenceIndicators().deleteWithResponse("myRg", "myWorkspace",
            "d9cd6f0b-96b9-3984-17cd-a779d1e15a93", com.azure.core.util.Context.NONE);
    }
}

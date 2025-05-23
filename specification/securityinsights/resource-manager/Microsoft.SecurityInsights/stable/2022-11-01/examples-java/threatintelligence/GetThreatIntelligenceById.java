
/**
 * Samples for ThreatIntelligenceIndicator Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2022-11-01/examples/
     * threatintelligence/GetThreatIntelligenceById.json
     */
    /**
     * Sample code: View a threat intelligence indicator by name.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void viewAThreatIntelligenceIndicatorByName(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.threatIntelligenceIndicators().getWithResponse("myRg", "myWorkspace",
            "e16ef847-962e-d7b6-9c8b-a33e4bd30e47", com.azure.core.util.Context.NONE);
    }
}

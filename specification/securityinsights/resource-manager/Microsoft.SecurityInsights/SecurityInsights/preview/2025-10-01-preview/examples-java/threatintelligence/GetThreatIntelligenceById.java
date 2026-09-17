
/**
 * Samples for ThreatIntelligenceIndicator Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/threatintelligence/GetThreatIntelligenceById.json
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

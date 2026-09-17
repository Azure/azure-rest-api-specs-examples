
/**
 * Samples for TriggeredAnalyticsRuleRunOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/triggeredAnalyticsRuleRuns/triggeredAnalyticsRuleRun_Get.json
     */
    /**
     * Sample code: triggeredAnalyticsRuleRun_Get.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        triggeredAnalyticsRuleRunGet(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.triggeredAnalyticsRuleRunOperations().getWithResponse("myRg", "myWorkspace",
            "65360bb0-8986-4ade-a89d-af3cf44d28aa", com.azure.core.util.Context.NONE);
    }
}

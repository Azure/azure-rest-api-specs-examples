
/**
 * Samples for GetTriggeredAnalyticsRuleRuns List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/triggeredAnalyticsRuleRuns/triggeredAnalyticsRuleRuns_Get.json
     */
    /**
     * Sample code: triggeredAnalyticsRuleRuns_Get.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        triggeredAnalyticsRuleRunsGet(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.getTriggeredAnalyticsRuleRuns().list("myRg", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}

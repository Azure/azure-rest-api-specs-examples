
/**
 * Samples for AlertRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/alertRules/GetFusionAlertRule.json
     */
    /**
     * Sample code: Get a Fusion alert rule.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAFusionAlertRule(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.alertRules().getWithResponse("myRg", "myWorkspace", "myFirstFusionRule",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for AlertRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2022-11-01/examples/alertRules/
     * GetScheduledAlertRule.json
     */
    /**
     * Sample code: Get a Scheduled alert rule.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAScheduledAlertRule(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.alertRules().getWithResponse("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
            com.azure.core.util.Context.NONE);
    }
}

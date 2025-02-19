
/**
 * Samples for Actions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2022-11-01/examples/actions/
     * GetActionOfAlertRuleById.json
     */
    /**
     * Sample code: Get an action of alert rule.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAnActionOfAlertRule(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.actions().getWithResponse("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
            "912bec42-cb66-4c03-ac63-1761b6898c3e", com.azure.core.util.Context.NONE);
    }
}

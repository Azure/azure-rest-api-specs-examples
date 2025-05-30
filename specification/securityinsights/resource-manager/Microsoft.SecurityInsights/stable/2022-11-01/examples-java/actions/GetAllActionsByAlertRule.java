
/**
 * Samples for Actions ListByAlertRule.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2022-11-01/examples/actions/
     * GetAllActionsByAlertRule.json
     */
    /**
     * Sample code: Get all actions of alert rule.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAllActionsOfAlertRule(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.actions().listByAlertRule("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
            com.azure.core.util.Context.NONE);
    }
}

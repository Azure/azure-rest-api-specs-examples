
/**
 * Samples for AutomationRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2022-11-01/examples/
     * automationRules/AutomationRules_Delete.json
     */
    /**
     * Sample code: AutomationRules_Delete.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        automationRulesDelete(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.automationRules().deleteWithResponse("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
            com.azure.core.util.Context.NONE);
    }
}

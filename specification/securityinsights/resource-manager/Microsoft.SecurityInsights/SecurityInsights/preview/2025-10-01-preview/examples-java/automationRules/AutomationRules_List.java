
/**
 * Samples for AutomationRules List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/automationRules/AutomationRules_List.json
     */
    /**
     * Sample code: AutomationRules_List.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void automationRulesList(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.automationRules().list("myRg", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}

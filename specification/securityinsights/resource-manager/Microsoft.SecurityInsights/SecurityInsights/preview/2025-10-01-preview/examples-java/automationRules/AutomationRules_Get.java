
/**
 * Samples for AutomationRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/automationRules/AutomationRules_Get.json
     */
    /**
     * Sample code: AutomationRules_Get.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void automationRulesGet(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.automationRules().getWithResponse("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
            com.azure.core.util.Context.NONE);
    }
}

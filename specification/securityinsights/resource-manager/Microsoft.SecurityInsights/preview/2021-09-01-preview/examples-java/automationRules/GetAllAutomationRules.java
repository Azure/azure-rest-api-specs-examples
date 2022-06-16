import com.azure.core.util.Context;

/** Samples for AutomationRules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/automationRules/GetAllAutomationRules.json
     */
    /**
     * Sample code: Get all automation rules.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllAutomationRules(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.automationRules().list("myRg", "myWorkspace", Context.NONE);
    }
}

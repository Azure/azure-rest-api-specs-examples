import com.azure.core.util.Context;

/** Samples for AlertRuleTemplates Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/alertRuleTemplates/GetAlertRuleTemplateById.json
     */
    /**
     * Sample code: Get alert rule template by Id.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAlertRuleTemplateById(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .alertRuleTemplates()
            .getWithResponse("myRg", "myWorkspace", "65360bb0-8986-4ade-a89d-af3cf44d28aa", Context.NONE);
    }
}

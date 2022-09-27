import com.azure.core.util.Context;

/** Samples for AlertRuleTemplates List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRuleTemplates/GetAlertRuleTemplates.json
     */
    /**
     * Sample code: Get all alert rule templates.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllAlertRuleTemplates(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.alertRuleTemplates().list("myRg", "myWorkspace", Context.NONE);
    }
}

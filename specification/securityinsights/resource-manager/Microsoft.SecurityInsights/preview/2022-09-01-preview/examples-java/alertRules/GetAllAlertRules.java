
/**
 * Samples for AlertRules List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * alertRules/GetAllAlertRules.json
     */
    /**
     * Sample code: Get all alert rules.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllAlertRules(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.alertRules().list("myRg", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for AlertRuleIncidents Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/getAlertRuleIncident.json
     */
    /**
     * Sample code: Get a single alert rule incident.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getASingleAlertRuleIncident(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getAlertRuleIncidents().getWithResponse("Rac46PostSwapRG",
            "myRuleName", "Website_started", com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for AlertRuleIncidents ListByAlertRule.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/listAlertRuleIncidents.json
     */
    /**
     * Sample code: List alert rule incidents.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAlertRuleIncidents(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getAlertRuleIncidents().listByAlertRule("Rac46PostSwapRG",
            "myRuleName", com.azure.core.util.Context.NONE);
    }
}

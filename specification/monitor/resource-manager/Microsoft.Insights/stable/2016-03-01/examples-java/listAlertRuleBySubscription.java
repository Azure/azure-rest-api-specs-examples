
/**
 * Samples for AlertRules List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/listAlertRuleBySubscription.
     * json
     */
    /**
     * Sample code: List alert rules.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAlertRules(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getAlertRules().list(com.azure.core.util.Context.NONE);
    }
}

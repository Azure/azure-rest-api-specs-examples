
/**
 * Samples for DataCollectionRules List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/preview/2021-09-01-preview/examples/
     * DataCollectionRulesListBySubscription.json
     */
    /**
     * Sample code: List data collection rules by subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDataCollectionRulesBySubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getDataCollectionRules()
            .list(com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for DataCollectionRuleAssociations ListByRule.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/preview/2021-09-01-preview/examples/
     * DataCollectionRuleAssociationsListByRule.json
     */
    /**
     * Sample code: List associations for specified data collection rule.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listAssociationsForSpecifiedDataCollectionRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getDataCollectionRuleAssociations()
            .listByRule("myResourceGroup", "myCollectionRule");
    }
}

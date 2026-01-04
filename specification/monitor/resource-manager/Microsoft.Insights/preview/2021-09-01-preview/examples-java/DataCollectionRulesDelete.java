
/**
 * Samples for DataCollectionRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/preview/2021-09-01-preview/examples/
     * DataCollectionRulesDelete.json
     */
    /**
     * Sample code: Delete data collection rule.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteDataCollectionRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getDataCollectionRules().delete("myResourceGroup",
            "myCollectionRule");
    }
}

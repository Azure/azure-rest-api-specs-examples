
import com.azure.core.util.Context;

/** Samples for DataCollectionRules GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/preview/2021-09-01-preview/examples/
     * DataCollectionRulesGet.json
     */
    /**
     * Sample code: Get data collection rule.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDataCollectionRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getDataCollectionRules()
            .getByResourceGroupWithResponse("myResourceGroup", "myCollectionRule", Context.NONE);
    }
}

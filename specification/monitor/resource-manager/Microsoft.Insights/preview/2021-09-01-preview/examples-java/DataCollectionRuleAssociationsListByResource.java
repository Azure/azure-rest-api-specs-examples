
import com.azure.core.util.Context;

/** Samples for DataCollectionRuleAssociations ListByResource. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/preview/2021-09-01-preview/examples/
     * DataCollectionRuleAssociationsListByResource.json
     */
    /**
     * Sample code: List associations for specified resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAssociationsForSpecifiedResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getDataCollectionRuleAssociations().listByResource(
            "subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVm",
            Context.NONE);
    }
}

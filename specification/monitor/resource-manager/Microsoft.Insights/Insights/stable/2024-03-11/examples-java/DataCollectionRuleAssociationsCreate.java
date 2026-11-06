
import com.azure.resourcemanager.monitor.fluent.models.DataCollectionRuleAssociationProxyOnlyResourceInner;

/**
 * Samples for DataCollectionRuleAssociations Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-11/DataCollectionRuleAssociationsCreate.json
     */
    /**
     * Sample code: Create or update association.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void createOrUpdateAssociation(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getDataCollectionRuleAssociations().createWithResponse(
            "subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVm",
            "myAssociation",
            new DataCollectionRuleAssociationProxyOnlyResourceInner().withDataCollectionRuleId(
                "/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Insights/dataCollectionRules/myCollectionRule"),
            com.azure.core.util.Context.NONE);
    }
}

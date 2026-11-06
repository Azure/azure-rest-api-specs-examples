
/**
 * Samples for DataCollectionRuleAssociations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-11/DataCollectionRuleAssociationsDelete.json
     */
    /**
     * Sample code: Delete association.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void deleteAssociation(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getDataCollectionRuleAssociations().deleteWithResponse(
            "subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVm",
            "myAssociation", com.azure.core.util.Context.NONE);
    }
}

/** Samples for Accounts ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_ListByResourceGroup.json
     */
    /**
     * Sample code: Accounts_ListByResourceGroup.
     *
     * @param manager Entry point to PurviewManager.
     */
    public static void accountsListByResourceGroup(com.azure.resourcemanager.purview.PurviewManager manager) {
        manager.accounts().listByResourceGroup("SampleResourceGroup", null, com.azure.core.util.Context.NONE);
    }
}

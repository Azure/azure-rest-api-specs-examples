/** Samples for Accounts GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_Get.json
     */
    /**
     * Sample code: Accounts_Get.
     *
     * @param manager Entry point to PurviewManager.
     */
    public static void accountsGet(com.azure.resourcemanager.purview.PurviewManager manager) {
        manager
            .accounts()
            .getByResourceGroupWithResponse("SampleResourceGroup", "account1", com.azure.core.util.Context.NONE);
    }
}

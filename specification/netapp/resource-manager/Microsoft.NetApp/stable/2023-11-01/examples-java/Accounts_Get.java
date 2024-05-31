
/**
 * Samples for Accounts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-11-01/examples/Accounts_Get.json
     */
    /**
     * Sample code: Accounts_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void accountsGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.accounts().getByResourceGroupWithResponse("myRG", "account1", com.azure.core.util.Context.NONE);
    }
}

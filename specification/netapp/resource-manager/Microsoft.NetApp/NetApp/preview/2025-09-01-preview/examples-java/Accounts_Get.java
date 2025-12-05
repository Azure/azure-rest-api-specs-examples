
/**
 * Samples for Accounts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Accounts_Get.json
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

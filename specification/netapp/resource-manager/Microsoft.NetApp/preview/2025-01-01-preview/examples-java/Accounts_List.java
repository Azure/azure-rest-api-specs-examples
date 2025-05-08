
/**
 * Samples for Accounts ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-01-01-preview/examples/Accounts_List.json
     */
    /**
     * Sample code: Accounts_List.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void accountsList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.accounts().listByResourceGroup("myRG", com.azure.core.util.Context.NONE);
    }
}

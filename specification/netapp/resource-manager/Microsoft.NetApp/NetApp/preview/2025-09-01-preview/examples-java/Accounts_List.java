
/**
 * Samples for Accounts ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Accounts_List.json
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

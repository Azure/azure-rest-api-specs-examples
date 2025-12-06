
/**
 * Samples for Accounts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Accounts_Delete.json
     */
    /**
     * Sample code: Accounts_Delete.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void accountsDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.accounts().delete("myRG", "account1", com.azure.core.util.Context.NONE);
    }
}

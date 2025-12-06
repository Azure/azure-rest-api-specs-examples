
/**
 * Samples for Accounts RenewCredentials.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Accounts_RenewCredentials.json
     */
    /**
     * Sample code: Accounts_RenewCredentials.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void accountsRenewCredentials(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.accounts().renewCredentials("myRG", "account1", com.azure.core.util.Context.NONE);
    }
}

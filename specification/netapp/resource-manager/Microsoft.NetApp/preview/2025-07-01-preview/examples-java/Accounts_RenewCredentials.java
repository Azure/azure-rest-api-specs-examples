
/**
 * Samples for Accounts RenewCredentials.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-07-01-preview/examples/
     * Accounts_RenewCredentials.json
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

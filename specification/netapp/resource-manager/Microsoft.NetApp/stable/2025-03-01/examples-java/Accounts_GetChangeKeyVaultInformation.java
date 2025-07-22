
/**
 * Samples for Accounts GetChangeKeyVaultInformation.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/
     * Accounts_GetChangeKeyVaultInformation.json
     */
    /**
     * Sample code: Accounts_GetChangeKeyVaultInformation.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void
        accountsGetChangeKeyVaultInformation(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.accounts().getChangeKeyVaultInformation("myRG", "account1", com.azure.core.util.Context.NONE);
    }
}

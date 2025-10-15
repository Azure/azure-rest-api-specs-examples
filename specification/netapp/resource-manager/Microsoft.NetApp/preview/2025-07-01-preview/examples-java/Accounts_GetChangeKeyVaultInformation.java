
/**
 * Samples for Accounts GetChangeKeyVaultInformation.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-07-01-preview/examples/
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

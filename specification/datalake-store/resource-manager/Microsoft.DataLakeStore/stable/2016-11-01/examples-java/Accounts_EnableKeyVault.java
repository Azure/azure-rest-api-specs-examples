/** Samples for Accounts EnableKeyVault. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_EnableKeyVault.json
     */
    /**
     * Sample code: Attempts to enable a user managed Key Vault for encryption of the specified Data Lake Store account.
     *
     * @param manager Entry point to DataLakeStoreManager.
     */
    public static void attemptsToEnableAUserManagedKeyVaultForEncryptionOfTheSpecifiedDataLakeStoreAccount(
        com.azure.resourcemanager.datalakestore.DataLakeStoreManager manager) {
        manager.accounts().enableKeyVaultWithResponse("contosorg", "contosoadla", com.azure.core.util.Context.NONE);
    }
}

/** Samples for Accounts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_Delete.json
     */
    /**
     * Sample code: Deletes the specified Data Lake Store account.
     *
     * @param manager Entry point to DataLakeStoreManager.
     */
    public static void deletesTheSpecifiedDataLakeStoreAccount(
        com.azure.resourcemanager.datalakestore.DataLakeStoreManager manager) {
        manager.accounts().delete("contosorg", "contosoadla", com.azure.core.util.Context.NONE);
    }
}

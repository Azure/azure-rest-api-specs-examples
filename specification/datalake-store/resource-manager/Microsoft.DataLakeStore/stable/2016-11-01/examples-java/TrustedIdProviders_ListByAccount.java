/** Samples for TrustedIdProviders ListByAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/TrustedIdProviders_ListByAccount.json
     */
    /**
     * Sample code: Lists the Data Lake Store trusted identity providers within the specified Data Lake Store account.
     *
     * @param manager Entry point to DataLakeStoreManager.
     */
    public static void listsTheDataLakeStoreTrustedIdentityProvidersWithinTheSpecifiedDataLakeStoreAccount(
        com.azure.resourcemanager.datalakestore.DataLakeStoreManager manager) {
        manager.trustedIdProviders().listByAccount("contosorg", "contosoadla", com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for TrustedIdProviders Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/
     * TrustedIdProviders_Delete.json
     */
    /**
     * Sample code: Deletes the specified trusted identity provider from the specified Data Lake Store account.
     * 
     * @param manager Entry point to DataLakeStoreManager.
     */
    public static void deletesTheSpecifiedTrustedIdentityProviderFromTheSpecifiedDataLakeStoreAccount(
        com.azure.resourcemanager.datalakestore.DataLakeStoreManager manager) {
        manager.trustedIdProviders().deleteWithResponse("contosorg", "contosoadla", "test_trusted_id_provider_name",
            com.azure.core.util.Context.NONE);
    }
}

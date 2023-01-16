/** Samples for TrustedIdProviders Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/TrustedIdProviders_Get.json
     */
    /**
     * Sample code: Gets the specified Data Lake Store trusted identity provider.
     *
     * @param manager Entry point to DataLakeStoreManager.
     */
    public static void getsTheSpecifiedDataLakeStoreTrustedIdentityProvider(
        com.azure.resourcemanager.datalakestore.DataLakeStoreManager manager) {
        manager
            .trustedIdProviders()
            .getWithResponse(
                "contosorg", "contosoadla", "test_trusted_id_provider_name", com.azure.core.util.Context.NONE);
    }
}

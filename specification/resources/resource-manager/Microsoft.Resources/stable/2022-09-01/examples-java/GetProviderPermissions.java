
/** Samples for Providers ProviderPermissions. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Resources/stable/2022-09-01/examples/GetProviderPermissions.
     * json
     */
    /**
     * Sample code: Get provider resource types.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getProviderResourceTypes(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().serviceClient().getProviders()
            .providerPermissionsWithResponse("Microsoft.TestRP", com.azure.core.util.Context.NONE);
    }
}

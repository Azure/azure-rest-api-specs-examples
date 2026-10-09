
/**
 * Samples for Providers ProviderPermissions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01/GetProviderPermissions.json
     */
    /**
     * Sample code: Get provider resource types.
     * 
     * @param manager Entry point to ResourceManager.
     */
    public static void getProviderResourceTypes(com.azure.resourcemanager.resources.ResourceManager manager) {
        manager.serviceClient().getProviders().providerPermissionsWithResponse("Microsoft.TestRP",
            com.azure.core.util.Context.NONE);
    }
}

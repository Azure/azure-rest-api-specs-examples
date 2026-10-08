
/**
 * Samples for Providers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01/GetProvider.json
     */
    /**
     * Sample code: Get provider.
     * 
     * @param manager Entry point to ResourceManager.
     */
    public static void getProvider(com.azure.resourcemanager.resources.ResourceManager manager) {
        manager.serviceClient().getProviders().getWithResponse("Microsoft.TestRP1", null,
            com.azure.core.util.Context.NONE);
    }
}

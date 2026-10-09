
/**
 * Samples for Providers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01/GetProviders.json
     */
    /**
     * Sample code: Get providers.
     * 
     * @param manager Entry point to ResourceManager.
     */
    public static void getProviders(com.azure.resourcemanager.resources.ResourceManager manager) {
        manager.serviceClient().getProviders().list(null, com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for PrivateLinkResources ListSupported.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ListSupportedPrivateLinkResources.json
     */
    /**
     * Sample code: ListSupportedPrivateLinkResources.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void
        listSupportedPrivateLinkResources(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getPrivateLinkResources().listSupported("rg1", "mysearchservice",
            com.azure.core.util.Context.NONE);
    }
}

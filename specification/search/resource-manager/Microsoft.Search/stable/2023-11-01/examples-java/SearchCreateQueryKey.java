
/**
 * Samples for QueryKeys Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/search/resource-manager/Microsoft.Search/stable/2023-11-01/examples/SearchCreateQueryKey.json
     */
    /**
     * Sample code: SearchCreateQueryKey.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void searchCreateQueryKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.searchServices().manager().serviceClient().getQueryKeys().createWithResponse("rg1", "mysearchservice",
            "Query key for browser-based clients", null, com.azure.core.util.Context.NONE);
    }
}

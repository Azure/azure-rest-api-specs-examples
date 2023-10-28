/** Samples for QueryKeys ListBySearchService. */
public final class Main {
    /*
     * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2023-11-01/examples/SearchListQueryKeysBySearchService.json
     */
    /**
     * Sample code: SearchListQueryKeysBySearchService.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void searchListQueryKeysBySearchService(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .searchServices()
            .manager()
            .serviceClient()
            .getQueryKeys()
            .listBySearchService("rg1", "mysearchservice", null, com.azure.core.util.Context.NONE);
    }
}

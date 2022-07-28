import com.azure.core.util.Context;

/** Samples for SharedPrivateLinkResources ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/ListSharedPrivateLinkResourcesByService.json
     */
    /**
     * Sample code: ListSharedPrivateLinkResourcesByService.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSharedPrivateLinkResourcesByService(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .searchServices()
            .manager()
            .serviceClient()
            .getSharedPrivateLinkResources()
            .listByService("rg1", "mysearchservice", null, Context.NONE);
    }
}

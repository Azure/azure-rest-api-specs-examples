
/**
 * Samples for SharedPrivateLinkResources Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/search/resource-manager/Microsoft.Search/stable/2023-11-01/examples/DeleteSharedPrivateLinkResource
     * .json
     */
    /**
     * Sample code: SharedPrivateLinkResourceDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void sharedPrivateLinkResourceDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.searchServices().manager().serviceClient().getSharedPrivateLinkResources().delete("rg1",
            "mysearchservice", "testResource", null, com.azure.core.util.Context.NONE);
    }
}

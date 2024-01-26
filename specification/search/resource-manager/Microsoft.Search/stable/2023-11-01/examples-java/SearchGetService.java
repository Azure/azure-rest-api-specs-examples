
/** Samples for Services GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/search/resource-manager/Microsoft.Search/stable/2023-11-01/examples/SearchGetService.json
     */
    /**
     * Sample code: SearchGetService.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void searchGetService(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.searchServices().manager().serviceClient().getServices().getByResourceGroupWithResponse("rg1",
            "mysearchservice", null, com.azure.core.util.Context.NONE);
    }
}

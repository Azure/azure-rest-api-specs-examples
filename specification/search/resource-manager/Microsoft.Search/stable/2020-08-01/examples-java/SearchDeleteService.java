import com.azure.core.util.Context;

/** Samples for Services Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/SearchDeleteService.json
     */
    /**
     * Sample code: SearchDeleteService.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void searchDeleteService(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .searchServices()
            .manager()
            .serviceClient()
            .getServices()
            .deleteWithResponse("rg1", "mysearchservice", null, Context.NONE);
    }
}

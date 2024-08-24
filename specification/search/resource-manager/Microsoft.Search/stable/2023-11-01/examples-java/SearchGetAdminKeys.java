
/**
 * Samples for AdminKeys Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/search/resource-manager/Microsoft.Search/stable/2023-11-01/examples/SearchGetAdminKeys.json
     */
    /**
     * Sample code: SearchGetAdminKeys.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void searchGetAdminKeys(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.searchServices().manager().serviceClient().getAdminKeys().getWithResponse("rg1", "mysearchservice", null,
            com.azure.core.util.Context.NONE);
    }
}

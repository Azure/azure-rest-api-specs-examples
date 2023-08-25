/** Samples for ProductWikisOperation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListProductWikis.json
     */
    /**
     * Sample code: ApiManagementGetApiWiki.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetApiWiki(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .productWikisOperations()
            .list(
                "rg1", "apimService1", "57d1f7558aa04f15146d9d8a", null, null, null, com.azure.core.util.Context.NONE);
    }
}

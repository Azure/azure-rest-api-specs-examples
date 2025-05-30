
/**
 * Samples for ProductWiki Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetProductWiki.json
     */
    /**
     * Sample code: ApiManagementGetProductWiki.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetProductWiki(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.productWikis().getWithResponse("rg1", "apimService1", "57d1f7558aa04f15146d9d8a",
            com.azure.core.util.Context.NONE);
    }
}

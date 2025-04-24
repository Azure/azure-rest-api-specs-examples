
/**
 * Samples for ProductApiLink CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateProductApiLink.json
     */
    /**
     * Sample code: ApiManagementCreateProductApiLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateProductApiLink(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.productApiLinks().define("link1").withExistingProduct("rg1", "apimService1", "testproduct").withApiId(
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/echo-api")
            .create();
    }
}

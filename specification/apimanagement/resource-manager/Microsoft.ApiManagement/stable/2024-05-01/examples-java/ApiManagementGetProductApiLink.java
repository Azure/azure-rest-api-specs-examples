
/**
 * Samples for ProductApiLink Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetProductApiLink.json
     */
    /**
     * Sample code: ApiManagementGetProductApiLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetProductApiLink(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.productApiLinks().getWithResponse("rg1", "apimService1", "testproduct", "link1",
            com.azure.core.util.Context.NONE);
    }
}

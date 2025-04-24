
/**
 * Samples for ProductGroupLink Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetProductGroupLink.json
     */
    /**
     * Sample code: ApiManagementGetProductGroupLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetProductGroupLink(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.productGroupLinks().getWithResponse("rg1", "apimService1", "testproduct", "link1",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for ProductGroupLink Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteProductGroupLink.json
     */
    /**
     * Sample code: ApiManagementDeleteProductGroupLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteProductGroupLink(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.productGroupLinks().deleteWithResponse("rg1", "apimService1", "testproduct", "link1",
            com.azure.core.util.Context.NONE);
    }
}

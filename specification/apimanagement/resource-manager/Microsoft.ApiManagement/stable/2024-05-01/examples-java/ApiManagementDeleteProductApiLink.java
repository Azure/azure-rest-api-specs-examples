
/**
 * Samples for ProductApiLink Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteProductApiLink.json
     */
    /**
     * Sample code: ApiManagementDeleteProductApiLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteProductApiLink(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.productApiLinks().deleteWithResponse("rg1", "apimService1", "testproduct", "link1",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for ProductGroupLink ListByProduct.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListProductGroupLinks.json
     */
    /**
     * Sample code: ApiManagementListProductGroupLinks.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListProductGroupLinks(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.productGroupLinks().listByProduct("rg1", "apimService1", "product1", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}

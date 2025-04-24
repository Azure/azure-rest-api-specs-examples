
/**
 * Samples for ProductApiLink ListByProduct.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListProductApiLinks.json
     */
    /**
     * Sample code: ApiManagementListProductApiLinks.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListProductApiLinks(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.productApiLinks().listByProduct("rg1", "apimService1", "product1", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}

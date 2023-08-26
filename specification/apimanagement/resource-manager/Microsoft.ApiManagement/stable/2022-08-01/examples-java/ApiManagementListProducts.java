/** Samples for Product ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListProducts.json
     */
    /**
     * Sample code: ApiManagementListProducts.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListProducts(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .products()
            .listByService("rg1", "apimService1", null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}

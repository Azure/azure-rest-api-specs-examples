/** Samples for Product GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadProduct.json
     */
    /**
     * Sample code: ApiManagementHeadProduct.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadProduct(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .products()
            .getEntityTagWithResponse("rg1", "apimService1", "unlimited", com.azure.core.util.Context.NONE);
    }
}

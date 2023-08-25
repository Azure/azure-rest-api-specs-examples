/** Samples for Product Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetProduct.json
     */
    /**
     * Sample code: ApiManagementGetProduct.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetProduct(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.products().getWithResponse("rg1", "apimService1", "unlimited", com.azure.core.util.Context.NONE);
    }
}

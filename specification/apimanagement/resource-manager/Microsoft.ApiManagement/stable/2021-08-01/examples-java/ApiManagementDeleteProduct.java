import com.azure.core.util.Context;

/** Samples for Product Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteProduct.json
     */
    /**
     * Sample code: ApiManagementDeleteProduct.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteProduct(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.products().deleteWithResponse("rg1", "apimService1", "testproduct", "*", true, Context.NONE);
    }
}

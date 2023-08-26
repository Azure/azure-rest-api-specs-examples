import com.azure.resourcemanager.apimanagement.models.ProductContract;

/** Samples for Product Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateProduct.json
     */
    /**
     * Sample code: ApiManagementUpdateProduct.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateProduct(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        ProductContract resource =
            manager
                .products()
                .getWithResponse("rg1", "apimService1", "testproduct", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withDisplayName("Test Template ProductName 4").withIfMatch("*").apply();
    }
}

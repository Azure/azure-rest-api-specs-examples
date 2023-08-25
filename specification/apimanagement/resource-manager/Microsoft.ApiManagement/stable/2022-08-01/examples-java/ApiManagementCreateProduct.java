/** Samples for Product CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateProduct.json
     */
    /**
     * Sample code: ApiManagementCreateProduct.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateProduct(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .products()
            .define("testproduct")
            .withExistingService("rg1", "apimService1")
            .withDisplayName("Test Template ProductName 4")
            .create();
    }
}

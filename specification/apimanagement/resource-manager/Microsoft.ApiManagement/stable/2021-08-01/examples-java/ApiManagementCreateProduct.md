```java
/** Samples for Product CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateProduct.json
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

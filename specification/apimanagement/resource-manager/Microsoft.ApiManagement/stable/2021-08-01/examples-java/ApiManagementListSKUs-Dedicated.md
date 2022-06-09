```java
import com.azure.core.util.Context;

/** Samples for ApiManagementServiceSkus ListAvailableServiceSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListSKUs-Dedicated.json
     */
    /**
     * Sample code: ApiManagementListSKUs-Dedicated.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListSKUsDedicated(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiManagementServiceSkus().listAvailableServiceSkus("rg1", "apimService1", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

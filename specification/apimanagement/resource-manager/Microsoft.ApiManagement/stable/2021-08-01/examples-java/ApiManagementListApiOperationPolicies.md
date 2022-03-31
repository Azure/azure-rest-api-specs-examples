Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ApiOperationPolicy ListByOperation. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListApiOperationPolicies.json
     */
    /**
     * Sample code: ApiManagementListApiOperationPolicies.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListApiOperationPolicies(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiOperationPolicies()
            .listByOperationWithResponse(
                "rg1", "apimService1", "599e2953193c3c0bd0b3e2fa", "599e29ab193c3c0bd0b3e2fb", Context.NONE);
    }
}
```

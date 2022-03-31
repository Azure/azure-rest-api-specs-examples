Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Gateway ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGatewayListKeys.json
     */
    /**
     * Sample code: ApiManagementGatewayListKeys.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGatewayListKeys(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.gateways().listKeysWithResponse("rg1", "apimService1", "gw1", Context.NONE);
    }
}
```

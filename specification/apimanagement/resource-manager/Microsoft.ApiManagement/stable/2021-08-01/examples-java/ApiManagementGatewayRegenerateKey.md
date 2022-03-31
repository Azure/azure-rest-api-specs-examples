Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.GatewayKeyRegenerationRequestContract;
import com.azure.resourcemanager.apimanagement.models.KeyType;

/** Samples for Gateway RegenerateKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGatewayRegenerateKey.json
     */
    /**
     * Sample code: ApiManagementGatewayRegenerateKey.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGatewayRegenerateKey(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .gateways()
            .regenerateKeyWithResponse(
                "rg1",
                "apimService1",
                "gwId",
                new GatewayKeyRegenerationRequestContract().withKeyType(KeyType.PRIMARY),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.AssociationContract;
import com.azure.resourcemanager.apimanagement.models.ProvisioningState;

/** Samples for GatewayApi CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateGatewayApi.json
     */
    /**
     * Sample code: ApiManagementCreateGatewayApi.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateGatewayApi(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .gatewayApis()
            .createOrUpdateWithResponse(
                "rg1",
                "apimService1",
                "gw1",
                "echo-api",
                new AssociationContract().withProvisioningState(ProvisioningState.CREATED),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateEndpoints GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PrivateEndpointGetForManualApproval.json
     */
    /**
     * Sample code: Get private endpoint with manual approval connection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getPrivateEndpointWithManualApprovalConnection(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getPrivateEndpoints()
            .getByResourceGroupWithResponse("rg1", "testPe", null, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.ConnectionResetSharedKeyInner;

/** Samples for VirtualNetworkGatewayConnections ResetSharedKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkGatewayConnectionResetSharedKey.json
     */
    /**
     * Sample code: ResetVirtualNetworkGatewayConnectionSharedKey.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void resetVirtualNetworkGatewayConnectionSharedKey(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworkGatewayConnections()
            .resetSharedKey("rg1", "conn1", new ConnectionResetSharedKeyInner().withKeyLength(128), Context.NONE);
    }
}
```

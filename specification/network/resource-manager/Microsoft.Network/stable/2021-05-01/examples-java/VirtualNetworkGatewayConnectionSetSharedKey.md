Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.ConnectionSharedKeyInner;

/** Samples for VirtualNetworkGatewayConnections SetSharedKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkGatewayConnectionSetSharedKey.json
     */
    /**
     * Sample code: SetVirtualNetworkGatewayConnectionSharedKey.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void setVirtualNetworkGatewayConnectionSharedKey(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworkGatewayConnections()
            .setSharedKey("rg1", "connS2S", new ConnectionSharedKeyInner().withValue("AzureAbc123"), Context.NONE);
    }
}
```

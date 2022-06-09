```java
import com.azure.core.util.Context;

/** Samples for VirtualNetworkGateways GetVpnclientConnectionHealth. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkGatewayGetVpnclientConnectionHealth.json
     */
    /**
     * Sample code: GetVirtualNetworkGatewayVpnclientConnectionHealth.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVirtualNetworkGatewayVpnclientConnectionHealth(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworkGateways()
            .getVpnclientConnectionHealth("p2s-vnet-test", "vpnp2sgw", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

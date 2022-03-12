Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.SubResource;
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkPeeringInner;

/** Samples for VirtualNetworkPeerings CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkPeeringCreate.json
     */
    /**
     * Sample code: Create peering.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createPeering(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworkPeerings()
            .createOrUpdate(
                "peerTest",
                "vnet1",
                "peer",
                new VirtualNetworkPeeringInner()
                    .withAllowVirtualNetworkAccess(true)
                    .withAllowForwardedTraffic(true)
                    .withAllowGatewayTransit(false)
                    .withUseRemoteGateways(false)
                    .withRemoteVirtualNetwork(
                        new SubResource()
                            .withId(
                                "/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2")),
                null,
                Context.NONE);
    }
}
```

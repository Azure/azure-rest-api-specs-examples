Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.VirtualRouterPeeringInner;

/** Samples for VirtualRouterPeerings CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualRouterPeeringPut.json
     */
    /**
     * Sample code: Create Virtual Router Peering.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createVirtualRouterPeering(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualRouterPeerings()
            .createOrUpdate(
                "rg1",
                "virtualRouter",
                "peering1",
                new VirtualRouterPeeringInner().withPeerAsn(20000L).withPeerIp("192.168.1.5"),
                Context.NONE);
    }
}
```

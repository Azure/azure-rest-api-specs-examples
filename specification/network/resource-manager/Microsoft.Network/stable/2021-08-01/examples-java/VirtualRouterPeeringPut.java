import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.VirtualRouterPeeringInner;

/** Samples for VirtualRouterPeerings CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualRouterPeeringPut.json
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

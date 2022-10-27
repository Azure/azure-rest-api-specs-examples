import com.azure.core.util.Context;

/** Samples for VirtualNetworkPeerings List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VirtualNetworkPeeringList.json
     */
    /**
     * Sample code: List peerings.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listPeerings(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkPeerings().list("peerTest", "vnet1", Context.NONE);
    }
}

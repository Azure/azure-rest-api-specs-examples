
/**
 * Samples for VirtualNetworks ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkList.json
     */
    /**
     * Sample code: List virtual networks in resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listVirtualNetworksInResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworks().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}

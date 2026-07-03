
/**
 * Samples for VirtualNetworkAppliances ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkAppliances_List.json
     */
    /**
     * Sample code: List virtual network appliances in resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listVirtualNetworkAppliancesInResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkAppliances().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}

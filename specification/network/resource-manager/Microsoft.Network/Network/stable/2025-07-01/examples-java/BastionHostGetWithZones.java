
/**
 * Samples for BastionHosts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/BastionHostGetWithZones.json
     */
    /**
     * Sample code: Get Bastion Host With Zones.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getBastionHostWithZones(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getBastionHosts().getByResourceGroupWithResponse("rg1", "bastionhosttenant",
            com.azure.core.util.Context.NONE);
    }
}

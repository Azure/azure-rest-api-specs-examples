
/**
 * Samples for IpamPools Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/IpamPools_UpdateAllocationBounds.json
     */
    /**
     * Sample code: Update the allocation size bounds on a Pool resource.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        updateTheAllocationSizeBoundsOnAPoolResource(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getIpamPools().updateWithResponse("rg1", "TestNetworkManager", "TestPool", null, null,
            com.azure.core.util.Context.NONE);
    }
}

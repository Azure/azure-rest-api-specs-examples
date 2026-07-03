
/**
 * Samples for IpAllocations GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/IpAllocationGet.json
     */
    /**
     * Sample code: Get IpAllocation.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getIpAllocation(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getIpAllocations().getByResourceGroupWithResponse("rg1", "test-ipallocation", null,
            com.azure.core.util.Context.NONE);
    }
}

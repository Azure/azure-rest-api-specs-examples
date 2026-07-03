
/**
 * Samples for IpAllocations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/IpAllocationDelete.json
     */
    /**
     * Sample code: Delete IpAllocation.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteIpAllocation(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getIpAllocations().delete("rg1", "test-ipallocation", com.azure.core.util.Context.NONE);
    }
}

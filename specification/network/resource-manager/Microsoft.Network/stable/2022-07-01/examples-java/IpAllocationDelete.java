import com.azure.core.util.Context;

/** Samples for IpAllocations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/IpAllocationDelete.json
     */
    /**
     * Sample code: Delete IpAllocation.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteIpAllocation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getIpAllocations().delete("rg1", "test-ipallocation", Context.NONE);
    }
}

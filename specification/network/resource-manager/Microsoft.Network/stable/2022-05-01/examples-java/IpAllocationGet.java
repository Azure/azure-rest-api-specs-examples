import com.azure.core.util.Context;

/** Samples for IpAllocations GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/IpAllocationGet.json
     */
    /**
     * Sample code: Get IpAllocation.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getIpAllocation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getIpAllocations()
            .getByResourceGroupWithResponse("rg1", "test-ipallocation", null, Context.NONE);
    }
}

import com.azure.core.util.Context;

/** Samples for VirtualHubs GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualHubGet.json
     */
    /**
     * Sample code: VirtualHubGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualHubGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualHubs()
            .getByResourceGroupWithResponse("rg1", "virtualHub1", Context.NONE);
    }
}

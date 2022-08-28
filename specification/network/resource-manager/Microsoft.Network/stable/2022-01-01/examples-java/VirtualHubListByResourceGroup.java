import com.azure.core.util.Context;

/** Samples for VirtualHubs ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VirtualHubListByResourceGroup.json
     */
    /**
     * Sample code: VirtualHubListByResourceGroup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualHubListByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualHubs().listByResourceGroup("rg1", Context.NONE);
    }
}

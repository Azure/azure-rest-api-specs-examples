
/**
 * Samples for NetworkVirtualApplianceConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkVirtualApplianceConnectionDelete.json
     */
    /**
     * Sample code: NetworkVirtualApplianceConnectionDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void networkVirtualApplianceConnectionDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkVirtualApplianceConnections().delete("rg1", "nva1",
            "connection1", com.azure.core.util.Context.NONE);
    }
}

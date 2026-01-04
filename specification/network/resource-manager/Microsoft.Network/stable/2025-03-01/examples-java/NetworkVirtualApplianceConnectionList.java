
/**
 * Samples for NetworkVirtualApplianceConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkVirtualApplianceConnectionList.json
     */
    /**
     * Sample code: NetworkVirtualApplianceConnectionList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void networkVirtualApplianceConnectionList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkVirtualApplianceConnections().list("rg1", "nva1",
            com.azure.core.util.Context.NONE);
    }
}

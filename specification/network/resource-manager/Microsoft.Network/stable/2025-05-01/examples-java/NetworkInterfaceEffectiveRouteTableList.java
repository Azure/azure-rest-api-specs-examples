
/**
 * Samples for NetworkInterfaces GetEffectiveRouteTable.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkInterfaceEffectiveRouteTableList.json
     */
    /**
     * Sample code: Show network interface effective route tables.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void showNetworkInterfaceEffectiveRouteTables(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkInterfaces().getEffectiveRouteTable("rg1", "nic1",
            com.azure.core.util.Context.NONE);
    }
}

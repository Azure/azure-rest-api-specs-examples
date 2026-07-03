
/**
 * Samples for NetworkInterfaces GetEffectiveRouteTable.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkInterfaceEffectiveRouteTableList.json
     */
    /**
     * Sample code: Show network interface effective route tables.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        showNetworkInterfaceEffectiveRouteTables(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkInterfaces().getEffectiveRouteTable("rg1", "nic1",
            com.azure.core.util.Context.NONE);
    }
}

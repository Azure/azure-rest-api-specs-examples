
/**
 * Samples for ExpressRouteCircuits ListArpTable.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitARPTableList.json
     */
    /**
     * Sample code: List ARP Table.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listARPTable(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuits().listArpTable("rg1", "circuitName", "peeringName",
            "devicePath", com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for ExpressRouteCircuits ListRoutesTable.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitRouteTableList.json
     */
    /**
     * Sample code: List Route Tables.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listRouteTables(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuits().listRoutesTable("rg1", "circuitName", "peeringName",
            "devicePath", com.azure.core.util.Context.NONE);
    }
}

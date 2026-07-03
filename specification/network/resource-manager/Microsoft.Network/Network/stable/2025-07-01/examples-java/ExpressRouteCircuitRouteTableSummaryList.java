
/**
 * Samples for ExpressRouteCircuits ListRoutesTableSummary.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitRouteTableSummaryList.json
     */
    /**
     * Sample code: List Route Table Summary.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listRouteTableSummary(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuits().listRoutesTableSummary("rg1", "circuitName", "peeringName",
            "devicePath", com.azure.core.util.Context.NONE);
    }
}

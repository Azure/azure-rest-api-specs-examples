import com.azure.core.util.Context;

/** Samples for ExpressRouteCircuits ListRoutesTableSummary. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRouteCircuitRouteTableSummaryList.json
     */
    /**
     * Sample code: List Route Table Summary.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listRouteTableSummary(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRouteCircuits()
            .listRoutesTableSummary("rg1", "circuitName", "peeringName", "devicePath", Context.NONE);
    }
}

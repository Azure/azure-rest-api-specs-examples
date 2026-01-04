
/**
 * Samples for ExpressRouteCircuits ListRoutesTable.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * ExpressRouteCircuitRouteTableList.json
     */
    /**
     * Sample code: List Route Tables.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listRouteTables(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteCircuits().listRoutesTable("rg1", "circuitName",
            "peeringName", "devicePath", com.azure.core.util.Context.NONE);
    }
}

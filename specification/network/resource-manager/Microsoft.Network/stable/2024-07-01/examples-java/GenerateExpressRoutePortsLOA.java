
import com.azure.resourcemanager.network.models.GenerateExpressRoutePortsLoaRequest;

/**
 * Samples for ExpressRoutePorts GenerateLoa.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/GenerateExpressRoutePortsLOA.
     * json
     */
    /**
     * Sample code: GenerateExpressRoutePortLOA.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void generateExpressRoutePortLOA(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRoutePorts().generateLoaWithResponse("rg1", "portName",
            new GenerateExpressRoutePortsLoaRequest().withCustomerName("customerName"),
            com.azure.core.util.Context.NONE);
    }
}

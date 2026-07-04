
import com.azure.resourcemanager.network.models.GenerateExpressRoutePortsLoaRequest;

/**
 * Samples for ExpressRoutePorts GenerateLoa.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/GenerateExpressRoutePortsLOA.json
     */
    /**
     * Sample code: GenerateExpressRoutePortLOA.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void generateExpressRoutePortLOA(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRoutePorts().generateLoaWithResponse("rg1", "portName",
            new GenerateExpressRoutePortsLoaRequest().withCustomerName("customerName"),
            com.azure.core.util.Context.NONE);
    }
}

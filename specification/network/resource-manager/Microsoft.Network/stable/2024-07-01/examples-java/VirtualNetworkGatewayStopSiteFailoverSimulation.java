
import com.azure.resourcemanager.network.models.ExpressRouteFailoverStopApiParameters;
import com.azure.resourcemanager.network.models.FailoverConnectionDetails;
import java.util.Arrays;

/**
 * Samples for VirtualNetworkGateways StopExpressRouteSiteFailoverSimulation.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * VirtualNetworkGatewayStopSiteFailoverSimulation.json
     */
    /**
     * Sample code: VirtualNetworkGatewayStopSiteFailoverSimulation.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualNetworkGatewayStopSiteFailoverSimulation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGateways().stopExpressRouteSiteFailoverSimulation(
            "rg1", "ergw",
            new ExpressRouteFailoverStopApiParameters().withPeeringLocation("Vancouver")
                .withWasSimulationSuccessful(true)
                .withDetails(Arrays.asList(
                    new FailoverConnectionDetails().withFailoverConnectionName("conn1").withFailoverLocation("Denver")
                        .withIsVerified(false),
                    new FailoverConnectionDetails().withFailoverConnectionName("conn2")
                        .withFailoverLocation("Amsterdam").withIsVerified(true))),
            com.azure.core.util.Context.NONE);
    }
}

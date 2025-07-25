
import com.azure.resourcemanager.network.models.GetInboundRoutesParameters;

/**
 * Samples for VirtualHubs GetInboundRoutes.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/GetInboundRoutes.json
     */
    /**
     * Sample code: Inbound Routes for the Virtual Hub on a Particular Connection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        inboundRoutesForTheVirtualHubOnAParticularConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualHubs().getInboundRoutes("rg1", "virtualHub1",
            new GetInboundRoutesParameters().withResourceUri(
                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteGateways/exrGw1/expressRouteConnections/exrConn1")
                .withConnectionType("ExpressRouteConnection"),
            com.azure.core.util.Context.NONE);
    }
}

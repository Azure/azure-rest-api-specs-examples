
import com.azure.resourcemanager.network.models.EffectiveRoutesParameters;

/**
 * Samples for VirtualHubs GetEffectiveVirtualHubRoutes.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * EffectiveRoutesListForConnection.json
     */
    /**
     * Sample code: Effective Routes for a Connection resource.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void effectiveRoutesForAConnectionResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualHubs().getEffectiveVirtualHubRoutes("rg1", "virtualHub1",
            new EffectiveRoutesParameters().withResourceId(
                "/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/expressRouteGatewayName/expressRouteConnections/connectionName")
                .withVirtualWanResourceType("ExpressRouteConnection"),
            com.azure.core.util.Context.NONE);
    }
}

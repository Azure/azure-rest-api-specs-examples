
import com.azure.resourcemanager.network.models.EffectiveRoutesParameters;

/**
 * Samples for VirtualHubs GetEffectiveVirtualHubRoutes.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * EffectiveRoutesListForRouteTable.json
     */
    /**
     * Sample code: Effective Routes for a Route Table resource.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void effectiveRoutesForARouteTableResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualHubs().getEffectiveVirtualHubRoutes("rg1", "virtualHub1",
            new EffectiveRoutesParameters().withResourceId(
                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1")
                .withVirtualWanResourceType("RouteTable"),
            com.azure.core.util.Context.NONE);
    }
}

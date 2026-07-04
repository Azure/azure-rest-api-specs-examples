
import com.azure.resourcemanager.network.models.EffectiveRoutesParameters;

/**
 * Samples for VirtualHubs GetEffectiveVirtualHubRoutes.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/EffectiveRoutesListForRouteTable.json
     */
    /**
     * Sample code: Effective Routes for a Route Table resource.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void effectiveRoutesForARouteTableResource(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualHubs().getEffectiveVirtualHubRoutes("rg1", "virtualHub1",
            new EffectiveRoutesParameters().withResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1")
                .withVirtualWanResourceType("RouteTable"),
            com.azure.core.util.Context.NONE);
    }
}

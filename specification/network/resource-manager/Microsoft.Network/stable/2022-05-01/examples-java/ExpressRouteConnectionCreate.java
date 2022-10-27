import com.azure.core.management.SubResource;
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.ExpressRouteConnectionInner;
import com.azure.resourcemanager.network.models.ExpressRouteCircuitPeeringId;
import com.azure.resourcemanager.network.models.PropagatedRouteTable;
import com.azure.resourcemanager.network.models.RoutingConfiguration;
import java.util.Arrays;

/** Samples for ExpressRouteConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/ExpressRouteConnectionCreate.json
     */
    /**
     * Sample code: ExpressRouteConnectionCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRouteConnectionCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRouteConnections()
            .createOrUpdate(
                "resourceGroupName",
                "gateway-2",
                "connectionName",
                new ExpressRouteConnectionInner()
                    .withId(
                        "/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/gateway-2/expressRouteConnections/connectionName")
                    .withName("connectionName")
                    .withExpressRouteCircuitPeering(
                        new ExpressRouteCircuitPeeringId()
                            .withId(
                                "/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteCircuits/circuitName/peerings/AzurePrivatePeering"))
                    .withAuthorizationKey("authorizationKey")
                    .withRoutingWeight(2)
                    .withRoutingConfiguration(
                        new RoutingConfiguration()
                            .withAssociatedRouteTable(
                                new SubResource()
                                    .withId(
                                        "/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable1"))
                            .withPropagatedRouteTables(
                                new PropagatedRouteTable()
                                    .withLabels(Arrays.asList("label1", "label2"))
                                    .withIds(
                                        Arrays
                                            .asList(
                                                new SubResource()
                                                    .withId(
                                                        "/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable1"),
                                                new SubResource()
                                                    .withId(
                                                        "/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable2"),
                                                new SubResource()
                                                    .withId(
                                                        "/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable3"))))
                            .withInboundRouteMap(
                                new SubResource()
                                    .withId(
                                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1"))
                            .withOutboundRouteMap(
                                new SubResource()
                                    .withId(
                                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap2"))),
                Context.NONE);
    }
}

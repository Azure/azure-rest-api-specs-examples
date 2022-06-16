import com.azure.core.management.SubResource;
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.HubVirtualNetworkConnectionInner;
import com.azure.resourcemanager.network.models.PropagatedRouteTable;
import com.azure.resourcemanager.network.models.RoutingConfiguration;
import com.azure.resourcemanager.network.models.StaticRoute;
import com.azure.resourcemanager.network.models.VnetRoute;
import java.util.Arrays;

/** Samples for HubVirtualNetworkConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/HubVirtualNetworkConnectionPut.json
     */
    /**
     * Sample code: HubVirtualNetworkConnectionPut.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void hubVirtualNetworkConnectionPut(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getHubVirtualNetworkConnections()
            .createOrUpdate(
                "rg1",
                "virtualHub1",
                "connection1",
                new HubVirtualNetworkConnectionInner()
                    .withRemoteVirtualNetwork(
                        new SubResource()
                            .withId(
                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/SpokeVnet1"))
                    .withEnableInternetSecurity(false)
                    .withRoutingConfiguration(
                        new RoutingConfiguration()
                            .withAssociatedRouteTable(
                                new SubResource()
                                    .withId(
                                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"))
                            .withPropagatedRouteTables(
                                new PropagatedRouteTable()
                                    .withLabels(Arrays.asList("label1", "label2"))
                                    .withIds(
                                        Arrays
                                            .asList(
                                                new SubResource()
                                                    .withId(
                                                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"))))
                            .withVnetRoutes(
                                new VnetRoute()
                                    .withStaticRoutes(
                                        Arrays
                                            .asList(
                                                new StaticRoute()
                                                    .withName("route1")
                                                    .withAddressPrefixes(Arrays.asList("10.1.0.0/16", "10.2.0.0/16"))
                                                    .withNextHopIpAddress("10.0.0.68"),
                                                new StaticRoute()
                                                    .withName("route2")
                                                    .withAddressPrefixes(Arrays.asList("10.3.0.0/16", "10.4.0.0/16"))
                                                    .withNextHopIpAddress("10.0.0.65"))))),
                Context.NONE);
    }
}

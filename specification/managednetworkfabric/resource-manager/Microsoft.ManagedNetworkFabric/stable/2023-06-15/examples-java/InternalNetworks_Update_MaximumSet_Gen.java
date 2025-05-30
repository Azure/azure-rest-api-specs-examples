
import com.azure.resourcemanager.managednetworkfabric.models.AllowASOverride;
import com.azure.resourcemanager.managednetworkfabric.models.BfdConfiguration;
import com.azure.resourcemanager.managednetworkfabric.models.BgpConfiguration;
import com.azure.resourcemanager.managednetworkfabric.models.BooleanEnumProperty;
import com.azure.resourcemanager.managednetworkfabric.models.ConnectedSubnet;
import com.azure.resourcemanager.managednetworkfabric.models.ExportRoutePolicy;
import com.azure.resourcemanager.managednetworkfabric.models.ImportRoutePolicy;
import com.azure.resourcemanager.managednetworkfabric.models.InternalNetwork;
import com.azure.resourcemanager.managednetworkfabric.models.IsMonitoringEnabled;
import com.azure.resourcemanager.managednetworkfabric.models.NeighborAddress;
import com.azure.resourcemanager.managednetworkfabric.models.StaticRouteConfiguration;
import com.azure.resourcemanager.managednetworkfabric.models.StaticRouteProperties;
import java.util.Arrays;

/**
 * Samples for InternalNetworks Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/
     * InternalNetworks_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: InternalNetworks_Update_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internalNetworksUpdateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        InternalNetwork resource = manager.internalNetworks().getWithResponse("example-rg", "example-l3domain",
            "example-internalnetwork", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withBgpConfiguration(new BgpConfiguration().withAnnotation("annotation")
                .withBfdConfiguration(new BfdConfiguration().withIntervalInMilliSeconds(300).withMultiplier(5))
                .withDefaultRouteOriginate(BooleanEnumProperty.TRUE).withAllowAS(10)
                .withAllowASOverride(AllowASOverride.ENABLE).withPeerAsn(61234L)
                .withIpv4ListenRangePrefixes(Arrays.asList("10.1.0.0/25"))
                .withIpv6ListenRangePrefixes(Arrays.asList("2fff::/66"))
                .withIpv4NeighborAddress(Arrays.asList(new NeighborAddress().withAddress("10.1.0.0")))
                .withIpv6NeighborAddress(Arrays.asList(new NeighborAddress().withAddress("2fff::"))))
            .withStaticRouteConfiguration(new StaticRouteConfiguration()
                .withBfdConfiguration(new BfdConfiguration().withIntervalInMilliSeconds(300).withMultiplier(15))
                .withIpv4Routes(Arrays.asList(
                    new StaticRouteProperties().withPrefix("20.20.20.20/25").withNextHop(Arrays.asList("10.0.0.1"))))
                .withIpv6Routes(Arrays
                    .asList(new StaticRouteProperties().withPrefix("2fff::/64").withNextHop(Arrays.asList("3ffe::1")))))
            .withMtu(1500)
            .withConnectedIPv4Subnets(
                Arrays.asList(new ConnectedSubnet().withAnnotation("annotation").withPrefix("10.0.0.0/24")))
            .withConnectedIPv6Subnets(
                Arrays.asList(new ConnectedSubnet().withAnnotation("annotation").withPrefix("3FFE:FFFF:0:CD30::a0/29")))
            .withImportRoutePolicyId(
                "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName")
            .withExportRoutePolicyId(
                "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName")
            .withImportRoutePolicy(new ImportRoutePolicy().withImportIpv4RoutePolicyId(
                "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName")
                .withImportIpv6RoutePolicyId(
                    "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"))
            .withExportRoutePolicy(new ExportRoutePolicy().withExportIpv4RoutePolicyId(
                "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName")
                .withExportIpv6RoutePolicyId(
                    "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"))
            .withIngressAclId(
                "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl")
            .withEgressAclId(
                "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl")
            .withIsMonitoringEnabled(IsMonitoringEnabled.TRUE).withAnnotation("annotation").apply();
    }
}

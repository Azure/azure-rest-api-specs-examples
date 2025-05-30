
import com.azure.resourcemanager.managednetworkfabric.models.AggregateRoute;
import com.azure.resourcemanager.managednetworkfabric.models.AggregateRouteConfiguration;
import com.azure.resourcemanager.managednetworkfabric.models.ConnectedSubnetRoutePolicy;
import com.azure.resourcemanager.managednetworkfabric.models.L3ExportRoutePolicy;
import com.azure.resourcemanager.managednetworkfabric.models.L3IsolationDomain;
import com.azure.resourcemanager.managednetworkfabric.models.RedistributeConnectedSubnets;
import com.azure.resourcemanager.managednetworkfabric.models.RedistributeStaticRoutes;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for L3IsolationDomains Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/
     * L3IsolationDomains_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: L3IsolationDomains_Update_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void l3IsolationDomainsUpdateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        L3IsolationDomain resource = manager.l3IsolationDomains()
            .getByResourceGroupWithResponse("example-rg", "example-l3domain", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("key4953", "fakeTokenPlaceholder"))
            .withRedistributeConnectedSubnets(RedistributeConnectedSubnets.TRUE)
            .withRedistributeStaticRoutes(RedistributeStaticRoutes.FALSE)
            .withAggregateRouteConfiguration(new AggregateRouteConfiguration()
                .withIpv4Routes(Arrays.asList(new AggregateRoute().withPrefix("10.0.0.0/24")))
                .withIpv6Routes(Arrays.asList(new AggregateRoute().withPrefix("3FFE:FFFF:0:CD30::a0/29"))))
            .withConnectedSubnetRoutePolicy(new ConnectedSubnetRoutePolicy().withExportRoutePolicyId(
                "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName")
                .withExportRoutePolicy(new L3ExportRoutePolicy().withExportIpv4RoutePolicyId(
                    "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy1")
                    .withExportIpv6RoutePolicyId(
                        "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy1")))
            .withAnnotation("annotation1").apply();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}

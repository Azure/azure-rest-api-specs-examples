
import com.azure.resourcemanager.managednetworkfabric.models.BfdConfiguration;
import com.azure.resourcemanager.managednetworkfabric.models.ExportRoutePolicy;
import com.azure.resourcemanager.managednetworkfabric.models.ExternalNetwork;
import com.azure.resourcemanager.managednetworkfabric.models.ExternalNetworkPatchPropertiesOptionAProperties;
import com.azure.resourcemanager.managednetworkfabric.models.ImportRoutePolicy;
import com.azure.resourcemanager.managednetworkfabric.models.L3OptionBProperties;
import com.azure.resourcemanager.managednetworkfabric.models.PeeringOption;
import com.azure.resourcemanager.managednetworkfabric.models.RouteTargetInformation;
import java.util.Arrays;

/**
 * Samples for ExternalNetworks Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/
     * ExternalNetworks_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExternalNetworks_Update_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void externalNetworksUpdateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        ExternalNetwork resource = manager.externalNetworks().getWithResponse("example-rg", "example-l3domain",
            "example-externalnetwork", com.azure.core.util.Context.NONE).getValue();
        resource.update().withPeeringOption(PeeringOption.OPTIONA)
            .withOptionBProperties(new L3OptionBProperties().withImportRouteTargets(Arrays.asList("65046:10039"))
                .withExportRouteTargets(Arrays.asList("65046:10039"))
                .withRouteTargets(new RouteTargetInformation().withImportIpv4RouteTargets(Arrays.asList("65046:10039"))
                    .withImportIpv6RouteTargets(Arrays.asList("65046:10039"))
                    .withExportIpv4RouteTargets(Arrays.asList("65046:10039"))
                    .withExportIpv6RouteTargets(Arrays.asList("65046:10039"))))
            .withOptionAProperties(new ExternalNetworkPatchPropertiesOptionAProperties()
                .withPrimaryIpv4Prefix("10.1.1.0/30").withPrimaryIpv6Prefix("3FFE:FFFF:0:CD30::a0/126")
                .withSecondaryIpv4Prefix("10.1.1.4/30").withSecondaryIpv6Prefix("3FFE:FFFF:0:CD30::a4/126")
                .withMtu(1500).withVlanId(1001).withPeerAsn(65047L)
                .withBfdConfiguration(new BfdConfiguration().withIntervalInMilliSeconds(300).withMultiplier(15))
                .withIngressAclId(
                    "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl")
                .withEgressAclId(
                    "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"))
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
            .withAnnotation("annotation1").apply();
    }
}

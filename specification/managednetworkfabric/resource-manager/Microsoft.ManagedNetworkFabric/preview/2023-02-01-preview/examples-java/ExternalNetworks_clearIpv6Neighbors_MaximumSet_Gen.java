import com.azure.resourcemanager.managednetworkfabric.models.EnableDisableOnResources;
import java.util.Arrays;

/** Samples for ExternalNetworks ClearIpv6Neighbors. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/ExternalNetworks_clearIpv6Neighbors_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExternalNetworks_clearIpv6Neighbors_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void externalNetworksClearIpv6NeighborsMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .externalNetworks()
            .clearIpv6Neighbors(
                "resourceGroupName",
                "example-l3domain",
                "example-externalnetwork",
                new EnableDisableOnResources()
                    .withResourceIds(
                        Arrays
                            .asList(
                                "/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/example-l3domain")),
                com.azure.core.util.Context.NONE);
    }
}

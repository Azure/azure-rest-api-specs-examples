import com.azure.resourcemanager.managednetworkfabric.models.EnableDisableOnResources;
import java.util.Arrays;

/** Samples for InternalNetworks ClearArpEntries. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/InternalNetworks_clearArpEntries_MaximumSet_Gen.json
     */
    /**
     * Sample code: InternalNetworks_clearArpEntries_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internalNetworksClearArpEntriesMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .internalNetworks()
            .clearArpEntries(
                "resourceGroupName",
                "example-l3domain",
                "example-internalnetwork",
                new EnableDisableOnResources()
                    .withResourceIds(
                        Arrays
                            .asList(
                                "/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/example-l3domain")),
                com.azure.core.util.Context.NONE);
    }
}

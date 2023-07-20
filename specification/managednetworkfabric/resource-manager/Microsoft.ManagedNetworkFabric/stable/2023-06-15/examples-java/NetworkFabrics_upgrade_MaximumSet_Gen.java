import com.azure.resourcemanager.managednetworkfabric.models.UpdateVersion;

/** Samples for NetworkFabrics Upgrade. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkFabrics_upgrade_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkFabrics_upgrade_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsUpgradeMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkFabrics()
            .upgrade(
                "example-rg",
                "example-fabric",
                new UpdateVersion().withVersion("version1"),
                com.azure.core.util.Context.NONE);
    }
}

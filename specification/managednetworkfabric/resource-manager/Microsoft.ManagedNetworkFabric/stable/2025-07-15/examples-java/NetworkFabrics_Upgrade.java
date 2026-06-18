
import com.azure.resourcemanager.managednetworkfabric.models.NetworkFabricUpgradeAction;
import com.azure.resourcemanager.managednetworkfabric.models.UpgradeNetworkFabricProperties;

/**
 * Samples for NetworkFabrics Upgrade.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_Upgrade.json
     */
    /**
     * Sample code: NetworkFabrics_Upgrade_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsUpgradeMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().upgrade("example-rg", "example-fabric",
            new UpgradeNetworkFabricProperties().withVersion("3.x.x").withAction(NetworkFabricUpgradeAction.START),
            com.azure.core.util.Context.NONE);
    }
}

import com.azure.resourcemanager.managednetworkfabric.models.UpdateVersion;

/** Samples for NetworkDevices Upgrade. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkDevices_Upgrade_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkDevices_Upgrade_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkDevicesUpgradeMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkDevices()
            .upgrade(
                "example-rg",
                "example-device",
                new UpdateVersion().withVersion("1.0.0"),
                com.azure.core.util.Context.NONE);
    }
}

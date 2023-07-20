/** Samples for NetworkDevices RefreshConfiguration. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkDevices_RefreshConfiguration_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkDevices_RefreshConfiguration_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkDevicesRefreshConfigurationMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkDevices().refreshConfiguration("example-rg", "example-device", com.azure.core.util.Context.NONE);
    }
}

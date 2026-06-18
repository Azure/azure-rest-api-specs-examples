
/**
 * Samples for NetworkDevices RefreshConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkDevices_RefreshConfiguration.json
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

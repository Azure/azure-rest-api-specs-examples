
/**
 * Samples for NetworkFabrics ViewDeviceConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_ViewDeviceConfiguration.json
     */
    /**
     * Sample code: NetworkFabrics_ViewDeviceConfiguration_MaximumSet.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsViewDeviceConfigurationMaximumSet(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().viewDeviceConfiguration("example-rg", "example-fabric",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for NetworkDevices Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkDevices_Delete.json
     */
    /**
     * Sample code: NetworkDevices_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkDevicesDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkDevices().delete("example-rg", "example-device", com.azure.core.util.Context.NONE);
    }
}

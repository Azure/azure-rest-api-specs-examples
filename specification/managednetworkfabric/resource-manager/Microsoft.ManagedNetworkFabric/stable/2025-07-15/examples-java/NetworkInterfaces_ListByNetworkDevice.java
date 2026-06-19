
/**
 * Samples for NetworkInterfaces ListByNetworkDevice.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkInterfaces_ListByNetworkDevice.json
     */
    /**
     * Sample code: NetworkInterfaces_ListByNetworkDevice_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkInterfacesListByNetworkDeviceMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkInterfaces().listByNetworkDevice("example-rg", "example-device",
            com.azure.core.util.Context.NONE);
    }
}

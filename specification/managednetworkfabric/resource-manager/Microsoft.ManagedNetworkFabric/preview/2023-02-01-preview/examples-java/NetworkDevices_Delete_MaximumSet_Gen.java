/** Samples for NetworkDevices Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkDevices_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkDevices_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkDevicesDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkDevices().delete("resourceGroupName", "networkDeviceName", com.azure.core.util.Context.NONE);
    }
}

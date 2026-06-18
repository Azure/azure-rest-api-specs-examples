
/**
 * Samples for NetworkDevices GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkDevices_Get.json
     */
    /**
     * Sample code: NetworkDevices_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkDevicesGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkDevices().getByResourceGroupWithResponse("example-rg", "example-device",
            com.azure.core.util.Context.NONE);
    }
}

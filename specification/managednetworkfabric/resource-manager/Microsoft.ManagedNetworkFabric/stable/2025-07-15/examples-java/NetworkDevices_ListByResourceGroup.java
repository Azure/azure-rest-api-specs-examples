
/**
 * Samples for NetworkDevices ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkDevices_ListByResourceGroup.json
     */
    /**
     * Sample code: NetworkDevices_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkDevicesListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkDevices().listByResourceGroup("example-rg", com.azure.core.util.Context.NONE);
    }
}

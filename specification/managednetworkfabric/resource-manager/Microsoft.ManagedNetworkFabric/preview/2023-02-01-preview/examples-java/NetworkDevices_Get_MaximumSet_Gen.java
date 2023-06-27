/** Samples for NetworkDevices GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkDevices_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkDevices_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkDevicesGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkDevices()
            .getByResourceGroupWithResponse("resourceGroupName", "networkDeviceName", com.azure.core.util.Context.NONE);
    }
}

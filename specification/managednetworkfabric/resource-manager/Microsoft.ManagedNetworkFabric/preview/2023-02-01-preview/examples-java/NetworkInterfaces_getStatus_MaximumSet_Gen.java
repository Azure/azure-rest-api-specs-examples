/** Samples for NetworkInterfaces GetStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkInterfaces_getStatus_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkInterfaces_getStatus_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkInterfacesGetStatusMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkInterfaces()
            .getStatus(
                "resourceGroupName", "networkDeviceName", "networkInterfaceName", com.azure.core.util.Context.NONE);
    }
}

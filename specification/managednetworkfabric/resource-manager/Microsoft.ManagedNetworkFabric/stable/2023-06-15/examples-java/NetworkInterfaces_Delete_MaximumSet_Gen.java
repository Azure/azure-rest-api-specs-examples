/** Samples for NetworkInterfaces Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkInterfaces_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkInterfaces_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkInterfacesDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkInterfaces().delete("rgNetworkDevices", "sjzd", "emrgu", com.azure.core.util.Context.NONE);
    }
}

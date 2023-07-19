/** Samples for NetworkInterfaces ListByNetworkDevice. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkInterfaces_ListByNetworkDevice_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkInterfaces_ListByNetworkDevice_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkInterfacesListByNetworkDeviceMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkInterfaces()
            .listByNetworkDevice("example-rg", "example-device", com.azure.core.util.Context.NONE);
    }
}

/** Samples for NetworkInterfaces Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkInterfaces_Create_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkInterfaces_Create_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkInterfacesCreateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkInterfaces()
            .define("example-interface")
            .withExistingNetworkDevice("example-rg", "example-device")
            .withAnnotation("annotation")
            .create();
    }
}

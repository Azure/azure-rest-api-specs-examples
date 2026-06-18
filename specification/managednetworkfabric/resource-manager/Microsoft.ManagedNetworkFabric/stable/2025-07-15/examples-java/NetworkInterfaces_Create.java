
/**
 * Samples for NetworkInterfaces Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkInterfaces_Create.json
     */
    /**
     * Sample code: NetworkInterfaces_Create_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkInterfacesCreateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkInterfaces().define("example-interface")
            .withExistingNetworkDevice("example-rg", "example-device").withAdditionalDescription("device 1")
            .withAnnotation("annotation").create();
    }
}

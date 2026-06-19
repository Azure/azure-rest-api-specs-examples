
/**
 * Samples for NetworkInterfaces Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkInterfaces_Get.json
     */
    /**
     * Sample code: NetworkInterfaces_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkInterfacesGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkInterfaces().getWithResponse("example-rg", "example-device", "example-interface",
            com.azure.core.util.Context.NONE);
    }
}

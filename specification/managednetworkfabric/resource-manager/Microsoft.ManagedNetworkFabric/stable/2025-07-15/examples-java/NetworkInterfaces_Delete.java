
/**
 * Samples for NetworkInterfaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkInterfaces_Delete.json
     */
    /**
     * Sample code: NetworkInterfaces_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkInterfacesDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkInterfaces().delete("example-rg", "example-device", "example-interface",
            com.azure.core.util.Context.NONE);
    }
}

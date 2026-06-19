
/**
 * Samples for NetworkFabrics ResyncPasswords.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_ResyncPasswords.json
     */
    /**
     * Sample code: Successful password resync for the Terminal Server and all Network Devices.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void successfulPasswordResyncForTheTerminalServerAndAllNetworkDevices(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().resyncPasswords("example-rg", "example-fabric", com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for NetworkDevices ResyncPasswords.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkDevices_ResyncPasswords_Error.json
     */
    /**
     * Sample code: Error while performing password resync.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void errorWhilePerformingPasswordResync(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkDevices().resyncPasswords("example-rg", "example-device", com.azure.core.util.Context.NONE);
    }
}

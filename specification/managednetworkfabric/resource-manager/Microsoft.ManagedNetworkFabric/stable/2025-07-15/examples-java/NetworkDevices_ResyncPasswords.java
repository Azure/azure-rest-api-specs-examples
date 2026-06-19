
/**
 * Samples for NetworkDevices ResyncPasswords.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkDevices_ResyncPasswords.json
     */
    /**
     * Sample code: Successful password resync.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void
        successfulPasswordResync(com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkDevices().resyncPasswords("example-rg", "example-device", com.azure.core.util.Context.NONE);
    }
}

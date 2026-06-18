
/**
 * Samples for NetworkDevices ResyncCertificates.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkDevices_ResyncCertificates_Error.json
     */
    /**
     * Sample code: Error while performing certificate resync.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void errorWhilePerformingCertificateResync(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkDevices().resyncCertificates("example-rg", "example-device", com.azure.core.util.Context.NONE);
    }
}

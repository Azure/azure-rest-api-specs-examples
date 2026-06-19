
/**
 * Samples for NetworkDevices ResyncCertificates.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkDevices_ResyncCertificates.json
     */
    /**
     * Sample code: Successful certificate resync.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void successfulCertificateResync(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkDevices().resyncCertificates("example-rg", "example-device", com.azure.core.util.Context.NONE);
    }
}

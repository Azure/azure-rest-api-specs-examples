
/**
 * Samples for NetworkFabrics ResyncCertificates.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_ResyncCertificates.json
     */
    /**
     * Sample code: Successful certificate resync for all Network Devices.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void successfulCertificateResyncForAllNetworkDevices(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().resyncCertificates("example-rg", "example-fabric", com.azure.core.util.Context.NONE);
    }
}

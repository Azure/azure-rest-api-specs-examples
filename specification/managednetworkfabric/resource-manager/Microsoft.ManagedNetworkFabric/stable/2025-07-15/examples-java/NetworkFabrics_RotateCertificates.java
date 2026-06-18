
/**
 * Samples for NetworkFabrics RotateCertificates.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_RotateCertificates.json
     */
    /**
     * Sample code: Successful certificate rotation for all Network Devices.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void successfulCertificateRotationForAllNetworkDevices(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().rotateCertificates("example-rg", "example-fabric", com.azure.core.util.Context.NONE);
    }
}

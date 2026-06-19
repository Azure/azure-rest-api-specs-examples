
/**
 * Samples for NetworkFabrics RotateCertificates.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_RotateCertificates_PartialSuccess.json
     */
    /**
     * Sample code: Partial failure to rotate certificates for some devices.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void partialFailureToRotateCertificatesForSomeDevices(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().rotateCertificates("example-rg", "example-fabric", com.azure.core.util.Context.NONE);
    }
}

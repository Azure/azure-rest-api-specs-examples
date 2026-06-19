
/**
 * Samples for NetworkFabrics RotateCertificates.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_RotateCertificates_Error.json
     */
    /**
     * Sample code: Total failure to rotate certificates.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void totalFailureToRotateCertificates(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().rotateCertificates("example-rg", "example-fabric", com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for NetworkFabrics ResyncCertificates.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_ResyncCertificates_Error.json
     */
    /**
     * Sample code: Total failure to resync certificates.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void totalFailureToResyncCertificates(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().resyncCertificates("example-rg", "example-fabric", com.azure.core.util.Context.NONE);
    }
}

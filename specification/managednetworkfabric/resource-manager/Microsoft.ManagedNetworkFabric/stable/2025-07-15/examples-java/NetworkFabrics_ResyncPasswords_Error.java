
/**
 * Samples for NetworkFabrics ResyncPasswords.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_ResyncPasswords_Error.json
     */
    /**
     * Sample code: Total failure to resync passwords.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void totalFailureToResyncPasswords(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().resyncPasswords("example-rg", "example-fabric", com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for NetworkFabrics ResyncPasswords.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_ResyncPasswords_PartialSuccess.json
     */
    /**
     * Sample code: Partial failure to resync passwords for some devices.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void partialFailureToResyncPasswordsForSomeDevices(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().resyncPasswords("example-rg", "example-fabric", com.azure.core.util.Context.NONE);
    }
}

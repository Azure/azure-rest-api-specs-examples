
/**
 * Samples for NetworkFabrics RotatePasswords.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_RotatePasswords_Error.json
     */
    /**
     * Sample code: Total failure to rotate passwords.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void totalFailureToRotatePasswords(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().rotatePasswords("example-rg", "example-fabric", com.azure.core.util.Context.NONE);
    }
}

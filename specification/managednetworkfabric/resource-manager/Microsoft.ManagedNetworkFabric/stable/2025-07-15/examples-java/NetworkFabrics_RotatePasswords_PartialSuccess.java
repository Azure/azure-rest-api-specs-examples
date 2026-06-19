
/**
 * Samples for NetworkFabrics RotatePasswords.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_RotatePasswords_PartialSuccess.json
     */
    /**
     * Sample code: Partial failure to rotate passwords for some devices.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void partialFailureToRotatePasswordsForSomeDevices(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().rotatePasswords("example-rg", "example-fabric", com.azure.core.util.Context.NONE);
    }
}

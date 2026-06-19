
/**
 * Samples for NetworkFabrics RotatePasswords.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_RotatePasswords.json
     */
    /**
     * Sample code: Successful password rotation for the Terminal Server and all Network Devices.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void successfulPasswordRotationForTheTerminalServerAndAllNetworkDevices(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().rotatePasswords("example-rg", "example-fabric", com.azure.core.util.Context.NONE);
    }
}

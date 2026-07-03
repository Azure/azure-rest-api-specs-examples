
/**
 * Samples for NetworkInterfaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkInterfaceDelete.json
     */
    /**
     * Sample code: Delete network interface.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteNetworkInterface(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkInterfaces().delete("rg1", "test-nic", com.azure.core.util.Context.NONE);
    }
}

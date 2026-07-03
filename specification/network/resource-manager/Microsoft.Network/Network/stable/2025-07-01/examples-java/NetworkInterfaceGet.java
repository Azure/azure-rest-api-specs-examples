
/**
 * Samples for NetworkInterfaces GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkInterfaceGet.json
     */
    /**
     * Sample code: Get network interface.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getNetworkInterface(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkInterfaces().getByResourceGroupWithResponse("rg1", "test-nic", null,
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for NetworkInterfaces GetCloudServiceNetworkInterface.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/CloudServiceNetworkInterfaceGet.json
     */
    /**
     * Sample code: Get cloud service network interface.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getCloudServiceNetworkInterface(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkInterfaces().getCloudServiceNetworkInterfaceWithResponse("rg1", "cs1",
            "TestVMRole_IN_0", "nic1", null, com.azure.core.util.Context.NONE);
    }
}

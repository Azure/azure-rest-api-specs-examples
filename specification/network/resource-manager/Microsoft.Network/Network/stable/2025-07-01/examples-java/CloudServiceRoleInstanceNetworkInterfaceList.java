
/**
 * Samples for NetworkInterfaces ListCloudServiceRoleInstanceNetworkInterfaces.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/CloudServiceRoleInstanceNetworkInterfaceList.json
     */
    /**
     * Sample code: List cloud service role instance network interfaces.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listCloudServiceRoleInstanceNetworkInterfaces(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkInterfaces().listCloudServiceRoleInstanceNetworkInterfaces("rg1", "cs1",
            "TestVMRole_IN_0", com.azure.core.util.Context.NONE);
    }
}

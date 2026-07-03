
/**
 * Samples for PublicIpAddresses ListCloudServiceRoleInstancePublicIpAddresses.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/CloudServiceRoleInstancePublicIpList.json
     */
    /**
     * Sample code: ListVMSSVMPublicIP.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listVMSSVMPublicIP(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPublicIpAddresses().listCloudServiceRoleInstancePublicIpAddresses("cs-tester", "cs1",
            "Test_VM_0", "nic1", "ip1", com.azure.core.util.Context.NONE);
    }
}

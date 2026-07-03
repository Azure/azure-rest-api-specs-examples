
/**
 * Samples for PublicIpAddresses GetCloudServicePublicIpAddress.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/CloudServicePublicIpGet.json
     */
    /**
     * Sample code: GetVMSSPublicIP.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getVMSSPublicIP(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPublicIpAddresses().getCloudServicePublicIpAddressWithResponse("cs-tester", "cs1",
            "Test_VM_0", "nic1", "ip1", "pub1", null, com.azure.core.util.Context.NONE);
    }
}

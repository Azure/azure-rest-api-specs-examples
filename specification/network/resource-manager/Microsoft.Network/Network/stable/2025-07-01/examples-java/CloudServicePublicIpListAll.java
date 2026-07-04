
/**
 * Samples for PublicIpAddresses ListCloudServicePublicIpAddresses.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/CloudServicePublicIpListAll.json
     */
    /**
     * Sample code: ListVMSSPublicIP.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listVMSSPublicIP(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPublicIpAddresses().listCloudServicePublicIpAddresses("cs-tester", "cs1",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for PublicIpAddresses GetCloudServicePublicIpAddress.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/CloudServicePublicIpGet.json
     */
    /**
     * Sample code: GetVMSSPublicIP.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVMSSPublicIP(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPublicIpAddresses().getCloudServicePublicIpAddressWithResponse(
            "cs-tester", "cs1", "Test_VM_0", "nic1", "ip1", "pub1", null, com.azure.core.util.Context.NONE);
    }
}

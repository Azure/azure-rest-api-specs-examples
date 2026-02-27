
/**
 * Samples for PublicIpAddresses ListCloudServicePublicIpAddresses.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/CloudServicePublicIpListAll.
     * json
     */
    /**
     * Sample code: ListVMSSPublicIP.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listVMSSPublicIP(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPublicIpAddresses().listCloudServicePublicIpAddresses("cs-tester",
            "cs1", com.azure.core.util.Context.NONE);
    }
}


import com.azure.resourcemanager.network.models.IsRollback;
import com.azure.resourcemanager.network.models.ReserveCloudServicePublicIpAddressRequest;

/**
 * Samples for PublicIpAddresses ReserveCloudServicePublicIpAddress.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/PublicIpAddressReserve.json
     */
    /**
     * Sample code: Reserve public IP address.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void reservePublicIPAddress(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPublicIpAddresses().reserveCloudServicePublicIpAddress("rg1",
            "test-ip", new ReserveCloudServicePublicIpAddressRequest().withIsRollback(IsRollback.FALSE),
            com.azure.core.util.Context.NONE);
    }
}

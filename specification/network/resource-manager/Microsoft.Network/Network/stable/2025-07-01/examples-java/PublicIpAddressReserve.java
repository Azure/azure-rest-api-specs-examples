
import com.azure.resourcemanager.network.models.IsRollback;
import com.azure.resourcemanager.network.models.ReserveCloudServicePublicIpAddressRequest;

/**
 * Samples for PublicIpAddresses ReserveCloudServicePublicIpAddress.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PublicIpAddressReserve.json
     */
    /**
     * Sample code: Reserve public IP address.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void reservePublicIPAddress(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPublicIpAddresses().reserveCloudServicePublicIpAddress("rg1", "test-ip",
            new ReserveCloudServicePublicIpAddressRequest().withIsRollback(IsRollback.FALSE),
            com.azure.core.util.Context.NONE);
    }
}

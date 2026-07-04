
/**
 * Samples for VirtualNetworks CheckIpAddressAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkCheckIPAddressAvailability.json
     */
    /**
     * Sample code: Check IP address availability.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void checkIPAddressAvailability(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworks().checkIpAddressAvailabilityWithResponse("rg1", "test-vnet",
            "10.0.1.4", com.azure.core.util.Context.NONE);
    }
}

import com.azure.core.util.Context;

/** Samples for VirtualNetworks CheckIpAddressAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VirtualNetworkCheckIPAddressAvailability.json
     */
    /**
     * Sample code: Check IP address availability.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void checkIPAddressAvailability(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworks()
            .checkIpAddressAvailabilityWithResponse("rg1", "test-vnet", "10.0.1.4", Context.NONE);
    }
}

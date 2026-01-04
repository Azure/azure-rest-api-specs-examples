
/**
 * Samples for PacketCaptures GetStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkWatcherPacketCaptureQueryStatus.json
     */
    /**
     * Sample code: Query packet capture status.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void queryPacketCaptureStatus(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPacketCaptures().getStatus("rg1", "nw1", "pc1",
            com.azure.core.util.Context.NONE);
    }
}

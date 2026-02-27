
/**
 * Samples for PacketCaptures Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkWatcherPacketCaptureStop.json
     */
    /**
     * Sample code: Stop packet capture.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void stopPacketCapture(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPacketCaptures().stop("rg1", "nw1", "pc1",
            com.azure.core.util.Context.NONE);
    }
}

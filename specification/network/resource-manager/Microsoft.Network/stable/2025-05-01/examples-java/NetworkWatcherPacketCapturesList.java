
/**
 * Samples for PacketCaptures List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkWatcherPacketCapturesList.json
     */
    /**
     * Sample code: List packet captures.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listPacketCaptures(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPacketCaptures().list("rg1", "nw1",
            com.azure.core.util.Context.NONE);
    }
}

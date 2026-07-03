
/**
 * Samples for PacketCaptures List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkWatcherPacketCapturesList.json
     */
    /**
     * Sample code: List packet captures.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listPacketCaptures(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPacketCaptures().list("rg1", "nw1", com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for PacketCaptures GetStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkWatcherPacketCaptureQueryStatus.json
     */
    /**
     * Sample code: Query packet capture status.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void queryPacketCaptureStatus(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPacketCaptures().getStatus("rg1", "nw1", "pc1", com.azure.core.util.Context.NONE);
    }
}

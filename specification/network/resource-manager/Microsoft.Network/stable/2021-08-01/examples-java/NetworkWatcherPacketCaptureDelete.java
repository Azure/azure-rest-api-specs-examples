import com.azure.core.util.Context;

/** Samples for PacketCaptures Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/NetworkWatcherPacketCaptureDelete.json
     */
    /**
     * Sample code: Delete packet capture.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletePacketCapture(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPacketCaptures().delete("rg1", "nw1", "pc1", Context.NONE);
    }
}

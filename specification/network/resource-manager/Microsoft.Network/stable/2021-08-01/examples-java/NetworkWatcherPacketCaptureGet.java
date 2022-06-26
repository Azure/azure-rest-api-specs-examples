import com.azure.core.util.Context;

/** Samples for PacketCaptures Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/NetworkWatcherPacketCaptureGet.json
     */
    /**
     * Sample code: Get packet capture.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getPacketCapture(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getPacketCaptures()
            .getWithResponse("rg1", "nw1", "pc1", Context.NONE);
    }
}

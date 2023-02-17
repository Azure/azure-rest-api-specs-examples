import com.azure.resourcemanager.mediaservices.models.EnabledProtocols;
import com.azure.resourcemanager.mediaservices.models.NoEncryption;

/** Samples for StreamingPolicies Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-policies-create-clear.json
     */
    /**
     * Sample code: Creates a Streaming Policy with clear streaming.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createsAStreamingPolicyWithClearStreaming(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .streamingPolicies()
            .define("clearStreamingPolicy")
            .withExistingMediaService("contosorg", "contosomedia")
            .withNoEncryption(
                new NoEncryption()
                    .withEnabledProtocols(
                        new EnabledProtocols()
                            .withDownload(true)
                            .withDash(true)
                            .withHls(true)
                            .withSmoothStreaming(true)))
            .create();
    }
}

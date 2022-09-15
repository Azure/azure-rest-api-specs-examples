import com.azure.core.util.Context;

/** Samples for StreamingPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/streaming-policy-get-by-name.json
     */
    /**
     * Sample code: Get a Streaming Policy by name.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getAStreamingPolicyByName(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.streamingPolicies().getWithResponse("contoso", "contosomedia", "clearStreamingPolicy", Context.NONE);
    }
}

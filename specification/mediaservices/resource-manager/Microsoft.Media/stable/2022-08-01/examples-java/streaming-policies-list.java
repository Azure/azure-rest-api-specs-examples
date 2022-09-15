import com.azure.core.util.Context;

/** Samples for StreamingPolicies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/streaming-policies-list.json
     */
    /**
     * Sample code: Lists Streaming Policies.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listsStreamingPolicies(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.streamingPolicies().list("contoso", "contosomedia", null, null, null, Context.NONE);
    }
}

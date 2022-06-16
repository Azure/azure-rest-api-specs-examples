import com.azure.core.util.Context;

/** Samples for StreamingPolicies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streaming-policies-delete.json
     */
    /**
     * Sample code: Delete a Streaming Policy.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void deleteAStreamingPolicy(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .streamingPolicies()
            .deleteWithResponse(
                "contoso", "contosomedia", "secureStreamingPolicyWithCommonEncryptionCbcsOnly", Context.NONE);
    }
}

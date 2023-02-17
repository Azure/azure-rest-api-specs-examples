/** Samples for StreamingPolicies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-policies-delete.json
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
                "contosorg",
                "contosomedia",
                "secureStreamingPolicyWithCommonEncryptionCbcsOnly",
                com.azure.core.util.Context.NONE);
    }
}

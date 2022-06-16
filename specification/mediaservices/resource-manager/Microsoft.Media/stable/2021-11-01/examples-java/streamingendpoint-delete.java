import com.azure.core.util.Context;

/** Samples for StreamingEndpoints Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streamingendpoint-delete.json
     */
    /**
     * Sample code: Delete a streaming endpoint.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void deleteAStreamingEndpoint(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.streamingEndpoints().delete("mediaresources", "slitestmedia10", "myStreamingEndpoint1", Context.NONE);
    }
}

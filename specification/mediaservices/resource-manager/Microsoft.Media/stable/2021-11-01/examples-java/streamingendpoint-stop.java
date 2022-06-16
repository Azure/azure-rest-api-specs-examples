import com.azure.core.util.Context;

/** Samples for StreamingEndpoints Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streamingendpoint-stop.json
     */
    /**
     * Sample code: Stop a streaming endpoint.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void stopAStreamingEndpoint(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.streamingEndpoints().stop("mediaresources", "slitestmedia10", "myStreamingEndpoint1", Context.NONE);
    }
}

/** Samples for StreamingEndpoints Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-11-01/examples/streamingendpoint-stop.json
     */
    /**
     * Sample code: Stop a streaming endpoint.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void stopAStreamingEndpoint(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .streamingEndpoints()
            .stop("mediaresources", "slitestmedia10", "myStreamingEndpoint1", com.azure.core.util.Context.NONE);
    }
}

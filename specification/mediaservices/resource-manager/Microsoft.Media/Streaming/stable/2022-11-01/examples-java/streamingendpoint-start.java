
/**
 * Samples for StreamingEndpoints Start.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-11-01/examples/
     * streamingendpoint-start.json
     */
    /**
     * Sample code: Start a streaming endpoint.
     * 
     * @param manager Entry point to MediaServicesManager.
     */
    public static void startAStreamingEndpoint(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.streamingEndpoints().start("mediaresources", "slitestmedia10", "myStreamingEndpoint1",
            com.azure.core.util.Context.NONE);
    }
}

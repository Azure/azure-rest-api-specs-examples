import com.azure.core.util.Context;

/** Samples for StreamingEndpoints Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streamingendpoint-start.json
     */
    /**
     * Sample code: Start a streaming endpoint.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void startAStreamingEndpoint(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.streamingEndpoints().start("mediaresources", "slitestmedia10", "myStreamingEndpoint1", Context.NONE);
    }
}

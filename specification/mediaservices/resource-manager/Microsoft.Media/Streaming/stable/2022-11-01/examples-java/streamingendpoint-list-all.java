/** Samples for StreamingEndpoints List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-11-01/examples/streamingendpoint-list-all.json
     */
    /**
     * Sample code: List all streaming endpoints.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listAllStreamingEndpoints(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.streamingEndpoints().list("mediaresources", "slitestmedia10", com.azure.core.util.Context.NONE);
    }
}

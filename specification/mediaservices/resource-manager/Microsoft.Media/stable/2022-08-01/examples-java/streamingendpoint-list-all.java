import com.azure.core.util.Context;

/** Samples for StreamingEndpoints List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/streamingendpoint-list-all.json
     */
    /**
     * Sample code: List all streaming endpoints.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listAllStreamingEndpoints(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.streamingEndpoints().list("mediaresources", "slitestmedia10", Context.NONE);
    }
}

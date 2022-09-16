import com.azure.core.util.Context;

/** Samples for StreamingEndpoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/streamingendpoint-list-by-name.json
     */
    /**
     * Sample code: Get a streaming endpoint by name.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getAStreamingEndpointByName(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .streamingEndpoints()
            .getWithResponse("mediaresources", "slitestmedia10", "myStreamingEndpoint1", Context.NONE);
    }
}

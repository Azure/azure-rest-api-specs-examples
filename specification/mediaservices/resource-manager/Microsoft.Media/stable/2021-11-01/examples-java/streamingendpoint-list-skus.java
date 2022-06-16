import com.azure.core.util.Context;

/** Samples for StreamingEndpoints Skus. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streamingendpoint-list-skus.json
     */
    /**
     * Sample code: List a streaming endpoint sku.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listAStreamingEndpointSku(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .streamingEndpoints()
            .skusWithResponse("mediaresources", "slitestmedia10", "myStreamingEndpoint1", Context.NONE);
    }
}

import com.azure.core.util.Context;

/** Samples for StreamingEndpoints OperationLocation. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/streamingendpoint-operation-location.json
     */
    /**
     * Sample code: Get the streaming endpoint operation status.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getTheStreamingEndpointOperationStatus(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .streamingEndpoints()
            .operationLocationWithResponse(
                "mediaresources",
                "slitestmedia10",
                "myStreamingEndpoint1",
                "62e4d893-d233-4005-988e-a428d9f77076",
                Context.NONE);
    }
}

import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/video-analyzer-private-endpoint-connection-delete.json
     */
    /**
     * Sample code: Delete private endpoint connection.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void deletePrivateEndpointConnection(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager
            .privateEndpointConnections()
            .deleteWithResponse("contoso", "contososports", "connectionName1", Context.NONE);
    }
}

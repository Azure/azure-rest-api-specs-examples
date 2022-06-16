import com.azure.core.util.Context;
import com.azure.resourcemanager.videoanalyzer.models.PrivateEndpointConnection;
import com.azure.resourcemanager.videoanalyzer.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.videoanalyzer.models.PrivateLinkServiceConnectionState;

/** Samples for PrivateEndpointConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/video-analyzer-private-endpoint-connection-put.json
     */
    /**
     * Sample code: Update private endpoint connection.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void updatePrivateEndpointConnection(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        PrivateEndpointConnection resource =
            manager
                .privateEndpointConnections()
                .getWithResponse("contoso", "contososports", "10000000-0000-0000-0000-000000000000", Context.NONE)
                .getValue();
        resource
            .update()
            .withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState()
                    .withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                    .withDescription("Test description."))
            .apply();
    }
}

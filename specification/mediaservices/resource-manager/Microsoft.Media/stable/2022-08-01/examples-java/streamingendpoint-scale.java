import com.azure.core.util.Context;
import com.azure.resourcemanager.mediaservices.models.StreamingEntityScaleUnit;

/** Samples for StreamingEndpoints Scale. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/streamingendpoint-scale.json
     */
    /**
     * Sample code: Scale a StreamingEndpoint.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void scaleAStreamingEndpoint(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .streamingEndpoints()
            .scale(
                "mediaresources",
                "slitestmedia10",
                "myStreamingEndpoint1",
                new StreamingEntityScaleUnit().withScaleUnit(5),
                Context.NONE);
    }
}

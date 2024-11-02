
import com.azure.resourcemanager.mediaservices.models.StreamingEntityScaleUnit;

/**
 * Samples for StreamingEndpoints Scale.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-11-01/examples/
     * streamingendpoint-scale.json
     */
    /**
     * Sample code: Scale a StreamingEndpoint.
     * 
     * @param manager Entry point to MediaServicesManager.
     */
    public static void scaleAStreamingEndpoint(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.streamingEndpoints().scale("mediaresources", "slitestmedia10", "myStreamingEndpoint1",
            new StreamingEntityScaleUnit().withScaleUnit(5), com.azure.core.util.Context.NONE);
    }
}

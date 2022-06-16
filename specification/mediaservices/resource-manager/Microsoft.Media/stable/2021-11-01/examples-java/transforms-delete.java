import com.azure.core.util.Context;

/** Samples for Transforms Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/transforms-delete.json
     */
    /**
     * Sample code: Delete a Transform.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void deleteATransform(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.transforms().deleteWithResponse("contosoresources", "contosomedia", "sampleTransform", Context.NONE);
    }
}

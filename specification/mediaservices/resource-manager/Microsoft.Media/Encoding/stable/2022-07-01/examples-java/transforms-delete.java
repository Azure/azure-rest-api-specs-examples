/** Samples for Transforms Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Encoding/stable/2022-07-01/examples/transforms-delete.json
     */
    /**
     * Sample code: Delete a Transform.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void deleteATransform(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .transforms()
            .deleteWithResponse(
                "contosoresources", "contosomedia", "sampleTransform", com.azure.core.util.Context.NONE);
    }
}

/** Samples for Transforms Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Encoding/stable/2022-07-01/examples/transforms-get-by-name.json
     */
    /**
     * Sample code: Get a Transform by name.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getATransformByName(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .transforms()
            .getWithResponse("contosoresources", "contosomedia", "sampleTransform", com.azure.core.util.Context.NONE);
    }
}

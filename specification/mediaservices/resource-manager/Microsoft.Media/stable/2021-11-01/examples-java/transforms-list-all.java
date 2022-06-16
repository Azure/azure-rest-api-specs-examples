import com.azure.core.util.Context;

/** Samples for Transforms List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/transforms-list-all.json
     */
    /**
     * Sample code: Lists the Transforms.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listsTheTransforms(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.transforms().list("contosoresources", "contosomedia", null, null, Context.NONE);
    }
}

import com.azure.core.util.Context;

/** Samples for Assets List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/assets-list-by-date.json
     */
    /**
     * Sample code: List Asset ordered by date.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listAssetOrderedByDate(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.assets().list("contoso", "contosomedia", null, null, "properties/created", Context.NONE);
    }
}

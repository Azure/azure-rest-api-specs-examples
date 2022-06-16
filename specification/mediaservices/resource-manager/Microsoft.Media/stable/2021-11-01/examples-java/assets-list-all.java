import com.azure.core.util.Context;

/** Samples for Assets List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/assets-list-all.json
     */
    /**
     * Sample code: List all Assets.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listAllAssets(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.assets().list("contoso", "contosomedia", null, null, null, Context.NONE);
    }
}

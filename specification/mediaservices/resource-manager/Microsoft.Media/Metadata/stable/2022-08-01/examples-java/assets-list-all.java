/** Samples for Assets List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/assets-list-all.json
     */
    /**
     * Sample code: List all Assets.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listAllAssets(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.assets().list("contoso", "contosomedia", null, null, null, com.azure.core.util.Context.NONE);
    }
}

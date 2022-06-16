import com.azure.core.util.Context;

/** Samples for Assets Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/assets-get-by-name.json
     */
    /**
     * Sample code: Get an Asset by name.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getAnAssetByName(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.assets().getWithResponse("contoso", "contosomedia", "ClimbingMountAdams", Context.NONE);
    }
}

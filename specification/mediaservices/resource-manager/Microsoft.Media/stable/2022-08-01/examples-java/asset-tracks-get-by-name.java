import com.azure.core.util.Context;

/** Samples for Tracks Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/asset-tracks-get-by-name.json
     */
    /**
     * Sample code: Get a Track by name.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getATrackByName(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.tracks().getWithResponse("contoso", "contosomedia", "ClimbingMountRainer", "text1", Context.NONE);
    }
}

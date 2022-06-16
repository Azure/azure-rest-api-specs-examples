import com.azure.core.util.Context;

/** Samples for Tracks List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/asset-tracks-list-all.json
     */
    /**
     * Sample code: Lists all Tracks.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listsAllTracks(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.tracks().list("contoso", "contosomedia", "ClimbingMountRainer", Context.NONE);
    }
}

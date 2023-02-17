/** Samples for Tracks List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/asset-tracks-list-all.json
     */
    /**
     * Sample code: Lists all Tracks.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listsAllTracks(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.tracks().list("contosorg", "contosomedia", "ClimbingMountRainer", com.azure.core.util.Context.NONE);
    }
}

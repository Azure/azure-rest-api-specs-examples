/** Samples for Tracks UpdateTrackData. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/asset-tracks-update-data.json
     */
    /**
     * Sample code: Update the data for a tracks.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void updateTheDataForATracks(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .tracks()
            .updateTrackData(
                "contoso", "contosomedia", "ClimbingMountRainer", "text2", com.azure.core.util.Context.NONE);
    }
}

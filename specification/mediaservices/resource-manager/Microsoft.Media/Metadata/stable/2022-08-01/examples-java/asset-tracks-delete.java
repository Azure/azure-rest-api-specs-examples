
/**
 * Samples for Tracks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/asset-tracks-
     * delete.json
     */
    /**
     * Sample code: Delete a Track.
     * 
     * @param manager Entry point to MediaServicesManager.
     */
    public static void deleteATrack(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.tracks().delete("contosorg", "contosomedia", "ClimbingMountRainer", "text2",
            com.azure.core.util.Context.NONE);
    }
}

import com.azure.core.util.Context;

/** Samples for Tracks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/asset-tracks-delete.json
     */
    /**
     * Sample code: Delete a Track.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void deleteATrack(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.tracks().delete("contoso", "contosomedia", "ClimbingMountRainer", "text2", Context.NONE);
    }
}

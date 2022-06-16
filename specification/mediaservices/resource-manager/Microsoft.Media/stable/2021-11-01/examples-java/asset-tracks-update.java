import com.azure.core.util.Context;
import com.azure.resourcemanager.mediaservices.models.AssetTrack;
import com.azure.resourcemanager.mediaservices.models.TextTrack;

/** Samples for Tracks Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/asset-tracks-update.json
     */
    /**
     * Sample code: Update a Track.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void updateATrack(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        AssetTrack resource =
            manager
                .tracks()
                .getWithResponse("contoso", "contosomedia", "ClimbingMountRainer", "text1", Context.NONE)
                .getValue();
        resource.update().withTrack(new TextTrack().withDisplayName("A new name")).apply();
    }
}

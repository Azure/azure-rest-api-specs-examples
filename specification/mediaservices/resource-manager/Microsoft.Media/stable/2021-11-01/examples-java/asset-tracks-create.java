import com.azure.resourcemanager.mediaservices.models.TextTrack;
import com.azure.resourcemanager.mediaservices.models.Visibility;

/** Samples for Tracks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/asset-tracks-create.json
     */
    /**
     * Sample code: Creates a Track.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createsATrack(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .tracks()
            .define("text3")
            .withExistingAsset("contoso", "contosomedia", "ClimbingMountRainer")
            .withTrack(
                new TextTrack()
                    .withFileName("text3.ttml")
                    .withDisplayName("A new track")
                    .withPlayerVisibility(Visibility.VISIBLE))
            .create();
    }
}

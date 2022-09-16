import com.azure.core.util.Context;
import com.azure.resourcemanager.mediaservices.models.AssetFilter;
import com.azure.resourcemanager.mediaservices.models.FirstQuality;
import com.azure.resourcemanager.mediaservices.models.PresentationTimeRange;

/** Samples for AssetFilters Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/assetFilters-update.json
     */
    /**
     * Sample code: Update an Asset Filter.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void updateAnAssetFilter(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        AssetFilter resource =
            manager
                .assetFilters()
                .getWithResponse(
                    "contoso", "contosomedia", "ClimbingMountRainer", "assetFilterWithTimeWindowAndTrack", Context.NONE)
                .getValue();
        resource
            .update()
            .withPresentationTimeRange(
                new PresentationTimeRange()
                    .withStartTimestamp(10L)
                    .withEndTimestamp(170000000L)
                    .withPresentationWindowDuration(9223372036854775000L)
                    .withLiveBackoffDuration(0L)
                    .withTimescale(10000000L)
                    .withForceEndTimestamp(false))
            .withFirstQuality(new FirstQuality().withBitrate(128000))
            .apply();
    }
}

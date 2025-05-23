
import com.azure.resourcemanager.mediaservices.models.FilterTrackPropertyCompareOperation;
import com.azure.resourcemanager.mediaservices.models.FilterTrackPropertyCondition;
import com.azure.resourcemanager.mediaservices.models.FilterTrackPropertyType;
import com.azure.resourcemanager.mediaservices.models.FilterTrackSelection;
import com.azure.resourcemanager.mediaservices.models.FirstQuality;
import com.azure.resourcemanager.mediaservices.models.PresentationTimeRange;
import java.util.Arrays;

/**
 * Samples for AccountFilters CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/accountFilters-
     * create.json
     */
    /**
     * Sample code: Create an Account Filter.
     * 
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createAnAccountFilter(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.accountFilters().define("newAccountFilter").withExistingMediaService("contosorg", "contosomedia")
            .withPresentationTimeRange(new PresentationTimeRange().withStartTimestamp(0L).withEndTimestamp(170000000L)
                .withPresentationWindowDuration(9223372036854775000L).withLiveBackoffDuration(0L)
                .withTimescale(10000000L).withForceEndTimestamp(false))
            .withFirstQuality(new FirstQuality().withBitrate(128000))
            .withTracks(Arrays.asList(
                new FilterTrackSelection().withTrackSelections(Arrays.asList(
                    new FilterTrackPropertyCondition().withProperty(FilterTrackPropertyType.TYPE).withValue("Audio")
                        .withOperation(FilterTrackPropertyCompareOperation.EQUAL),
                    new FilterTrackPropertyCondition().withProperty(FilterTrackPropertyType.LANGUAGE).withValue("en")
                        .withOperation(FilterTrackPropertyCompareOperation.NOT_EQUAL),
                    new FilterTrackPropertyCondition().withProperty(FilterTrackPropertyType.FOUR_CC).withValue("EC-3")
                        .withOperation(FilterTrackPropertyCompareOperation.NOT_EQUAL))),
                new FilterTrackSelection().withTrackSelections(Arrays.asList(
                    new FilterTrackPropertyCondition().withProperty(FilterTrackPropertyType.TYPE).withValue("Video")
                        .withOperation(FilterTrackPropertyCompareOperation.EQUAL),
                    new FilterTrackPropertyCondition().withProperty(FilterTrackPropertyType.BITRATE)
                        .withValue("3000000-5000000").withOperation(FilterTrackPropertyCompareOperation.EQUAL)))))
            .create();
    }
}

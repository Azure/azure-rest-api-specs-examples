Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_2.0.0/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.mediaservices.models.FilterTrackPropertyCompareOperation;
import com.azure.resourcemanager.mediaservices.models.FilterTrackPropertyCondition;
import com.azure.resourcemanager.mediaservices.models.FilterTrackPropertyType;
import com.azure.resourcemanager.mediaservices.models.FilterTrackSelection;
import com.azure.resourcemanager.mediaservices.models.FirstQuality;
import com.azure.resourcemanager.mediaservices.models.PresentationTimeRange;
import java.util.Arrays;

/** Samples for AssetFilters CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/assetFilters-create.json
     */
    /**
     * Sample code: Create an Asset Filter.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createAnAssetFilter(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .assetFilters()
            .define("newAssetFilter")
            .withExistingAsset("contoso", "contosomedia", "ClimbingMountRainer")
            .withPresentationTimeRange(
                new PresentationTimeRange()
                    .withStartTimestamp(0L)
                    .withEndTimestamp(170000000L)
                    .withPresentationWindowDuration(9223372036854775000L)
                    .withLiveBackoffDuration(0L)
                    .withTimescale(10000000L)
                    .withForceEndTimestamp(false))
            .withFirstQuality(new FirstQuality().withBitrate(128000))
            .withTracks(
                Arrays
                    .asList(
                        new FilterTrackSelection()
                            .withTrackSelections(
                                Arrays
                                    .asList(
                                        new FilterTrackPropertyCondition()
                                            .withProperty(FilterTrackPropertyType.TYPE)
                                            .withValue("Audio")
                                            .withOperation(FilterTrackPropertyCompareOperation.EQUAL),
                                        new FilterTrackPropertyCondition()
                                            .withProperty(FilterTrackPropertyType.LANGUAGE)
                                            .withValue("en")
                                            .withOperation(FilterTrackPropertyCompareOperation.NOT_EQUAL),
                                        new FilterTrackPropertyCondition()
                                            .withProperty(FilterTrackPropertyType.FOUR_CC)
                                            .withValue("EC-3")
                                            .withOperation(FilterTrackPropertyCompareOperation.NOT_EQUAL))),
                        new FilterTrackSelection()
                            .withTrackSelections(
                                Arrays
                                    .asList(
                                        new FilterTrackPropertyCondition()
                                            .withProperty(FilterTrackPropertyType.TYPE)
                                            .withValue("Video")
                                            .withOperation(FilterTrackPropertyCompareOperation.EQUAL),
                                        new FilterTrackPropertyCondition()
                                            .withProperty(FilterTrackPropertyType.BITRATE)
                                            .withValue("3000000-5000000")
                                            .withOperation(FilterTrackPropertyCompareOperation.EQUAL)))))
            .create();
    }
}
```

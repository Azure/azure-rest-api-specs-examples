import com.azure.core.util.Context;

/** Samples for AssetFilters Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/assetFilters-delete.json
     */
    /**
     * Sample code: Delete an Asset Filter.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void deleteAnAssetFilter(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .assetFilters()
            .deleteWithResponse(
                "contoso", "contosomedia", "ClimbingMountRainer", "assetFilterWithTimeWindowAndTrack", Context.NONE);
    }
}

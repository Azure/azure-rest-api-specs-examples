import com.azure.core.util.Context;

/** Samples for AssetFilters Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/assetFilters-get-by-name.json
     */
    /**
     * Sample code: Get an Asset Filter by name.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getAnAssetFilterByName(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .assetFilters()
            .getWithResponse(
                "contoso", "contosomedia", "ClimbingMountRainer", "assetFilterWithTimeWindowAndTrack", Context.NONE);
    }
}

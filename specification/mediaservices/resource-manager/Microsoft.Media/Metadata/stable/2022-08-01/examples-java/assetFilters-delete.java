/** Samples for AssetFilters Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/assetFilters-delete.json
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
                "contosorg",
                "contosomedia",
                "ClimbingMountRainer",
                "assetFilterWithTimeWindowAndTrack",
                com.azure.core.util.Context.NONE);
    }
}

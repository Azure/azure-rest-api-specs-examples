import com.azure.core.util.Context;

/** Samples for AssetFilters List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/assetFilters-list-all.json
     */
    /**
     * Sample code: List all Asset Filters.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listAllAssetFilters(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.assetFilters().list("contoso", "contosomedia", "ClimbingMountRainer", Context.NONE);
    }
}

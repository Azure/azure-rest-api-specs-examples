/** Samples for Assets List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/assets-list-in-date-range.json
     */
    /**
     * Sample code: List Asset created in a date range.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listAssetCreatedInADateRange(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .assets()
            .list(
                "contosorg",
                "contosomedia",
                "properties/created gt 2012-06-01 and properties/created lt 2013-07-01",
                null,
                "properties/created",
                com.azure.core.util.Context.NONE);
    }
}

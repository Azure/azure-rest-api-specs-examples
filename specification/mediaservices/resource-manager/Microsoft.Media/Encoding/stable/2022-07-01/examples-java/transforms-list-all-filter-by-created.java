/** Samples for Transforms List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Encoding/stable/2022-07-01/examples/transforms-list-all-filter-by-created.json
     */
    /**
     * Sample code: Lists the Transforms filter by created.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listsTheTransformsFilterByCreated(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .transforms()
            .list(
                "contosoresources",
                "contosomedia",
                "properties/created gt 2021-06-01T00:00:00.0000000Z and properties/created le"
                    + " 2021-06-01T00:00:10.0000000Z",
                "properties/created",
                com.azure.core.util.Context.NONE);
    }
}

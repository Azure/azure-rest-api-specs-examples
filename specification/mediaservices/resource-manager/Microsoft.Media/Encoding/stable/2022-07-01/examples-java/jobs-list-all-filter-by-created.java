/** Samples for Jobs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Encoding/stable/2022-07-01/examples/jobs-list-all-filter-by-created.json
     */
    /**
     * Sample code: Lists Jobs for the Transform filter by created.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listsJobsForTheTransformFilterByCreated(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .jobs()
            .list(
                "contosoresources",
                "contosomedia",
                "exampleTransform",
                "properties/created ge 2021-06-01T00:00:10.0000000Z and properties/created le"
                    + " 2021-06-01T00:00:20.0000000Z",
                "properties/created",
                com.azure.core.util.Context.NONE);
    }
}

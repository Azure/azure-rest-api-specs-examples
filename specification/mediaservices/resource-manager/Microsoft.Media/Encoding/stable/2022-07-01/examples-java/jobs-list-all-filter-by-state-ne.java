/** Samples for Jobs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Encoding/stable/2022-07-01/examples/jobs-list-all-filter-by-state-ne.json
     */
    /**
     * Sample code: Lists Jobs for the Transform filter by state not equal.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listsJobsForTheTransformFilterByStateNotEqual(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .jobs()
            .list(
                "contosoresources",
                "contosomedia",
                "exampleTransform",
                "properties/state ne Microsoft.Media.JobState'processing'",
                null,
                com.azure.core.util.Context.NONE);
    }
}

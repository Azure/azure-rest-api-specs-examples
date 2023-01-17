/** Samples for Jobs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Encoding/stable/2022-07-01/examples/jobs-list-all-filter-by-name-and-state.json
     */
    /**
     * Sample code: Lists Jobs for the Transform filter by name and state.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listsJobsForTheTransformFilterByNameAndState(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .jobs()
            .list(
                "contosoresources",
                "contosomedia",
                "exampleTransform",
                "name eq 'job3' and properties/state eq Microsoft.Media.JobState'finished'",
                null,
                com.azure.core.util.Context.NONE);
    }
}

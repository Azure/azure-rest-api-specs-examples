/** Samples for Jobs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Encoding/stable/2022-07-01/examples/jobs-list-all-filter-by-name.json
     */
    /**
     * Sample code: Lists Jobs for the Transform filter by name.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listsJobsForTheTransformFilterByName(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .jobs()
            .list(
                "contosoresources",
                "contosomedia",
                "exampleTransform",
                "name eq 'job1' or name eq 'job2'",
                "name",
                com.azure.core.util.Context.NONE);
    }
}

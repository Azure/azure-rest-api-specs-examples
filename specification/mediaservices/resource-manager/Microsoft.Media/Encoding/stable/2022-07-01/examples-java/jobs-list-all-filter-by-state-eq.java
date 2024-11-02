
/**
 * Samples for Jobs List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mediaservices/resource-manager/Microsoft.Media/Encoding/stable/2022-07-01/examples/jobs-list-all-
     * filter-by-state-eq.json
     */
    /**
     * Sample code: Lists Jobs for the Transform filter by state equal.
     * 
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listsJobsForTheTransformFilterByStateEqual(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.jobs().list("contosoresources", "contosomedia", "exampleTransform",
            "properties/state eq Microsoft.Media.JobState'Processing'", null, com.azure.core.util.Context.NONE);
    }
}

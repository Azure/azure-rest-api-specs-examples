import com.azure.core.util.Context;

/** Samples for Jobs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/jobs-list-all-filter-by-lastmodified.json
     */
    /**
     * Sample code: Lists Jobs for the Transform filter by lastmodified.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listsJobsForTheTransformFilterByLastmodified(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .jobs()
            .list(
                "contosoresources",
                "contosomedia",
                "exampleTransform",
                "properties/lastmodified ge 2021-11-01T00:00:10.0000000Z and properties/lastmodified le"
                    + " 2021-11-01T00:00:20.0000000Z",
                "properties/lastmodified desc",
                Context.NONE);
    }
}

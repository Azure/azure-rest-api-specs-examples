
/**
 * Samples for Jobs List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mediaservices/resource-manager/Microsoft.Media/Encoding/stable/2022-07-01/examples/jobs-list-all.
     * json
     */
    /**
     * Sample code: Lists all of the Jobs for the Transform.
     * 
     * @param manager Entry point to MediaServicesManager.
     */
    public static void
        listsAllOfTheJobsForTheTransform(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.jobs().list("contosoresources", "contosomedia", "exampleTransform", null, null,
            com.azure.core.util.Context.NONE);
    }
}

import com.azure.core.util.Context;

/** Samples for Jobs Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/jobs-delete.json
     */
    /**
     * Sample code: Delete a Job.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void deleteAJob(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .jobs()
            .deleteWithResponse("contosoresources", "contosomedia", "exampleTransform", "jobToDelete", Context.NONE);
    }
}

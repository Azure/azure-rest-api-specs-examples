/** Samples for Jobs CancelJob. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Encoding/stable/2022-07-01/examples/jobs-cancel.json
     */
    /**
     * Sample code: Cancel a Job.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void cancelAJob(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .jobs()
            .cancelJobWithResponse(
                "contosoresources", "contosomedia", "exampleTransform", "job1", com.azure.core.util.Context.NONE);
    }
}

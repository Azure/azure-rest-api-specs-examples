import com.azure.core.util.Context;

/** Samples for Jobs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/jobs-get-by-name.json
     */
    /**
     * Sample code: Get a Job by name.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getAJobByName(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.jobs().getWithResponse("contosoresources", "contosomedia", "exampleTransform", "job1", Context.NONE);
    }
}

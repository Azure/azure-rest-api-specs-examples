import com.azure.core.util.Context;
import com.azure.resourcemanager.mediaservices.models.Job;
import com.azure.resourcemanager.mediaservices.models.JobInputAsset;
import com.azure.resourcemanager.mediaservices.models.JobOutputAsset;
import com.azure.resourcemanager.mediaservices.models.Priority;
import java.util.Arrays;

/** Samples for Jobs Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/jobs-update.json
     */
    /**
     * Sample code: Update a Job.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void updateAJob(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        Job resource =
            manager
                .jobs()
                .getWithResponse("contosoresources", "contosomedia", "exampleTransform", "job1", Context.NONE)
                .getValue();
        resource
            .update()
            .withDescription("Example job to illustrate update.")
            .withInput(new JobInputAsset().withAssetName("job1-InputAsset"))
            .withOutputs(Arrays.asList(new JobOutputAsset().withAssetName("job1-OutputAsset")))
            .withPriority(Priority.HIGH)
            .apply();
    }
}

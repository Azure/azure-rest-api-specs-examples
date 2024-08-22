
import com.azure.resourcemanager.machinelearning.models.ComputeResource;
import com.azure.resourcemanager.machinelearning.models.ScaleSettings;
import com.azure.resourcemanager.machinelearning.models.ScaleSettingsInformation;
import java.time.Duration;

/**
 * Samples for Compute Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Compute/patch.json
     */
    /**
     * Sample code: Update a AmlCompute Compute.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        updateAAmlComputeCompute(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        ComputeResource resource = manager.computes()
            .getWithResponse("testrg123", "workspaces123", "compute123", com.azure.core.util.Context.NONE).getValue();
        resource
            .update().withProperties(new ScaleSettingsInformation().withScaleSettings(new ScaleSettings()
                .withMaxNodeCount(4).withMinNodeCount(4).withNodeIdleTimeBeforeScaleDown(Duration.parse("PT5M"))))
            .apply();
    }
}

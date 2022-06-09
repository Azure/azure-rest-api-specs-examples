```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.machinelearning.models.ComputeResource;
import com.azure.resourcemanager.machinelearning.models.ScaleSettings;
import com.azure.resourcemanager.machinelearning.models.ScaleSettingsInformation;
import java.time.Duration;

/** Samples for Compute Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Compute/patch.json
     */
    /**
     * Sample code: Update a AmlCompute Compute.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void updateAAmlComputeCompute(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        ComputeResource resource =
            manager.computes().getWithResponse("testrg123", "workspaces123", "compute123", Context.NONE).getValue();
        resource
            .update()
            .withProperties(
                new ScaleSettingsInformation()
                    .withScaleSettings(
                        new ScaleSettings()
                            .withMaxNodeCount(4)
                            .withMinNodeCount(4)
                            .withNodeIdleTimeBeforeScaleDown(Duration.parse("PT5M"))))
            .apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.1/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.

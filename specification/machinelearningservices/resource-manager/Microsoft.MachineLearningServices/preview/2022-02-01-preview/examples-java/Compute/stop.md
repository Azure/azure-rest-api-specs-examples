```java
import com.azure.core.util.Context;

/** Samples for Compute Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Compute/stop.json
     */
    /**
     * Sample code: Stop ComputeInstance Compute.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void stopComputeInstanceCompute(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.computes().stop("testrg123", "workspaces123", "compute123", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.1/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.

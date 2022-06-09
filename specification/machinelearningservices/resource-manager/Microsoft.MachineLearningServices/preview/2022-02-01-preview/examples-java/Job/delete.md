```java
import com.azure.core.util.Context;

/** Samples for Jobs Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/delete.json
     */
    /**
     * Sample code: Delete Job.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void deleteJob(com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.jobs().delete("test-rg", "my-aml-workspace", "string", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.1/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.

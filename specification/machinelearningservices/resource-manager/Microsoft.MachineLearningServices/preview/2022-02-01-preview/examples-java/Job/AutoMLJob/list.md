Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.2/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Jobs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/AutoMLJob/list.json
     */
    /**
     * Sample code: List AutoML Job.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listAutoMLJob(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.jobs().list("test-rg", "my-aml-workspace", null, null, null, null, null, null, Context.NONE);
    }
}
```

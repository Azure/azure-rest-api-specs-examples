Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.2/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Jobs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/CommandJob/list.json
     */
    /**
     * Sample code: List Command Job.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listCommandJob(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.jobs().list("test-rg", "my-aml-workspace", null, "string", "string", null, null, null, Context.NONE);
    }
}
```

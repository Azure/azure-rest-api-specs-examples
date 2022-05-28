Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.2/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for EnvironmentContainers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/EnvironmentContainer/list.json
     */
    /**
     * Sample code: List Environment Container.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listEnvironmentContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.environmentContainers().list("testrg123", "testworkspace", null, null, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.1/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for EnvironmentContainers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/EnvironmentContainer/delete.json
     */
    /**
     * Sample code: Delete Environment Container.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void deleteEnvironmentContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.environmentContainers().deleteWithResponse("testrg123", "testworkspace", "testContainer", Context.NONE);
    }
}
```

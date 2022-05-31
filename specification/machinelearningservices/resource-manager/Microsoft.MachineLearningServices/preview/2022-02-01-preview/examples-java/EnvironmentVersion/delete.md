Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.1/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for EnvironmentVersions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/EnvironmentVersion/delete.json
     */
    /**
     * Sample code: Delete Environment Version.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void deleteEnvironmentVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .environmentVersions()
            .deleteWithResponse("test-rg", "my-aml-workspace", "string", "string", Context.NONE);
    }
}
```

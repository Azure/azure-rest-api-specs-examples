Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.2/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ModelVersions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/ModelVersion/list.json
     */
    /**
     * Sample code: List Model Version.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listModelVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .modelVersions()
            .list(
                "test-rg",
                "my-aml-workspace",
                "string",
                null,
                "string",
                1,
                "string",
                "string",
                1,
                "string",
                "string",
                null,
                null,
                Context.NONE);
    }
}
```

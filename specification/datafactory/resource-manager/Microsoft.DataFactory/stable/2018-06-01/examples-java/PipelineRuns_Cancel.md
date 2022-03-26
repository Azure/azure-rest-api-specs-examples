Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-datafactory_1.0.0-beta.13/sdk/datafactory/azure-resourcemanager-datafactory/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PipelineRuns Cancel. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/PipelineRuns_Cancel.json
     */
    /**
     * Sample code: PipelineRuns_Cancel.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void pipelineRunsCancel(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager
            .pipelineRuns()
            .cancelWithResponse(
                "exampleResourceGroup",
                "exampleFactoryName",
                "16ac5348-ff82-4f95-a80d-638c1d47b721",
                null,
                Context.NONE);
    }
}
```

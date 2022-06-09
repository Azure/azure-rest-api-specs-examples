```java
import com.azure.core.util.Context;

/** Samples for PipelineRuns List. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-12-01-preview/examples/PipelineRunList.json
     */
    /**
     * Sample code: PipelineRunList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void pipelineRunList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getPipelineRuns()
            .list("myResourceGroup", "myRegistry", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ExportPipelines Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-12-01-preview/examples/ExportPipelineDelete.json
     */
    /**
     * Sample code: ExportPipelineDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void exportPipelineDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getExportPipelines()
            .delete("myResourceGroup", "myRegistry", "myExportPipeline", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ImportPipelines Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-12-01-preview/examples/ImportPipelineDelete.json
     */
    /**
     * Sample code: ImportPipelineDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void importPipelineDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getImportPipelines()
            .delete("myResourceGroup", "myRegistry", "myImportPipeline", Context.NONE);
    }
}
```

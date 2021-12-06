Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.containerregistry.fluent.models.PipelineRunInner;
import com.azure.resourcemanager.containerregistry.models.PipelineRunRequest;
import com.azure.resourcemanager.containerregistry.models.PipelineRunSourceProperties;
import com.azure.resourcemanager.containerregistry.models.PipelineRunSourceType;

/** Samples for PipelineRuns Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-12-01-preview/examples/PipelineRunCreate_Import.json
     */
    /**
     * Sample code: PipelineRunCreate_Import.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void pipelineRunCreateImport(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getPipelineRuns()
            .create(
                "myResourceGroup",
                "myRegistry",
                "myPipelineRun",
                new PipelineRunInner()
                    .withRequest(
                        new PipelineRunRequest()
                            .withPipelineResourceId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/importPipelines/myImportPipeline")
                            .withSource(
                                new PipelineRunSourceProperties()
                                    .withType(PipelineRunSourceType.AZURE_STORAGE_BLOB)
                                    .withName("myblob.tar.gz"))
                            .withCatalogDigest("sha256@"))
                    .withForceUpdateTag("2020-03-04T17:23:21.9261521+00:00"),
                Context.NONE);
    }
}
```

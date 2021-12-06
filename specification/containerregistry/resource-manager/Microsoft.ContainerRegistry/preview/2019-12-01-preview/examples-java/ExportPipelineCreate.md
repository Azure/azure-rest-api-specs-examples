Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.containerregistry.fluent.models.ExportPipelineInner;
import com.azure.resourcemanager.containerregistry.models.ExportPipelineTargetProperties;
import com.azure.resourcemanager.containerregistry.models.IdentityProperties;
import com.azure.resourcemanager.containerregistry.models.PipelineOptions;
import com.azure.resourcemanager.containerregistry.models.ResourceIdentityType;
import java.util.Arrays;

/** Samples for ExportPipelines Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-12-01-preview/examples/ExportPipelineCreate.json
     */
    /**
     * Sample code: ExportPipelineCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void exportPipelineCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getExportPipelines()
            .create(
                "myResourceGroup",
                "myRegistry",
                "myExportPipeline",
                new ExportPipelineInner()
                    .withLocation("westus")
                    .withIdentity(new IdentityProperties().withType(ResourceIdentityType.SYSTEM_ASSIGNED))
                    .withTarget(
                        new ExportPipelineTargetProperties()
                            .withType("AzureStorageBlobContainer")
                            .withUri("https://accountname.blob.core.windows.net/containername")
                            .withKeyVaultUri("https://myvault.vault.azure.net/secrets/acrexportsas"))
                    .withOptions(Arrays.asList(PipelineOptions.OVERWRITE_BLOBS)),
                Context.NONE);
    }
}
```

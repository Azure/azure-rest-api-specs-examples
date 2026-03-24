
import com.azure.resourcemanager.containerregistry.fluent.models.PipelineRunInner;
import com.azure.resourcemanager.containerregistry.models.PipelineRunRequest;
import com.azure.resourcemanager.containerregistry.models.PipelineRunSourceProperties;
import com.azure.resourcemanager.containerregistry.models.PipelineRunSourceType;

/**
 * Samples for PipelineRuns Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/PipelineRunCreate_Import.json
     */
    /**
     * Sample code: PipelineRunCreate_Import.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        pipelineRunCreateImport(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getPipelineRuns().create("myResourceGroup", "myRegistry", "myPipelineRun",
            new PipelineRunInner().withRequest(new PipelineRunRequest().withPipelineResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/importPipelines/myImportPipeline")
                .withSource(new PipelineRunSourceProperties().withType(PipelineRunSourceType.AZURE_STORAGE_BLOB)
                    .withName("myblob.tar.gz"))
                .withCatalogDigest("sha256@")).withForceUpdateTag("2020-03-04T17:23:21.9261521+00:00"),
            com.azure.core.util.Context.NONE);
    }
}


import com.azure.resourcemanager.containerregistry.fluent.models.PipelineRunInner;
import com.azure.resourcemanager.containerregistry.models.PipelineRunRequest;
import com.azure.resourcemanager.containerregistry.models.PipelineRunTargetProperties;
import com.azure.resourcemanager.containerregistry.models.PipelineRunTargetType;
import java.util.Arrays;

/**
 * Samples for PipelineRuns Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/PipelineRunCreate_Export.json
     */
    /**
     * Sample code: PipelineRunCreate_Export.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        pipelineRunCreateExport(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getPipelineRuns().create("myResourceGroup", "myRegistry", "myPipelineRun",
            new PipelineRunInner().withRequest(new PipelineRunRequest().withPipelineResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/exportPipelines/myExportPipeline")
                .withArtifacts(Arrays.asList("sourceRepository/hello-world",
                    "sourceRepository2@sha256:00000000000000000000000000000000000"))
                .withTarget(new PipelineRunTargetProperties().withType(PipelineRunTargetType.AZURE_STORAGE_BLOB)
                    .withName("myblob.tar.gz"))),
            com.azure.core.util.Context.NONE);
    }
}

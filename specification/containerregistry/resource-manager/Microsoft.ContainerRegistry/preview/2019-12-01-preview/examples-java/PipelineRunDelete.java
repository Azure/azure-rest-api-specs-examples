import com.azure.core.util.Context;

/** Samples for PipelineRuns Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-12-01-preview/examples/PipelineRunDelete.json
     */
    /**
     * Sample code: PipelineRunDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void pipelineRunDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getPipelineRuns()
            .delete("myResourceGroup", "myRegistry", "myPipelineRun", Context.NONE);
    }
}

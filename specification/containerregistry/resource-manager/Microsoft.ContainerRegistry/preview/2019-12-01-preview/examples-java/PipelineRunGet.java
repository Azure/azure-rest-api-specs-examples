import com.azure.core.util.Context;

/** Samples for PipelineRuns Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-12-01-preview/examples/PipelineRunGet.json
     */
    /**
     * Sample code: PipelineRunGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void pipelineRunGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getPipelineRuns()
            .getWithResponse("myResourceGroup", "myRegistry", "myPipelineRun", Context.NONE);
    }
}

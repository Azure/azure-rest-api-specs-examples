import com.azure.core.util.Context;

/** Samples for ImportPipelines Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-12-01-preview/examples/ImportPipelineGet.json
     */
    /**
     * Sample code: ImportPipelineGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void importPipelineGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getImportPipelines()
            .getWithResponse("myResourceGroup", "myRegistry", "myImportPipeline", Context.NONE);
    }
}

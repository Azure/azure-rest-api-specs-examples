import com.azure.core.util.Context;

/** Samples for ImportPipelines List. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-12-01-preview/examples/ImportPipelineList.json
     */
    /**
     * Sample code: ImportPipelineList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void importPipelineList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getImportPipelines()
            .list("myResourceGroup", "myRegistry", Context.NONE);
    }
}

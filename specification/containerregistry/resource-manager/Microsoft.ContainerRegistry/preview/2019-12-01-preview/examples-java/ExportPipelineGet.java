import com.azure.core.util.Context;

/** Samples for ExportPipelines Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-12-01-preview/examples/ExportPipelineGet.json
     */
    /**
     * Sample code: ExportPipelineGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void exportPipelineGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getExportPipelines()
            .getWithResponse("myResourceGroup", "myRegistry", "myExportPipeline", Context.NONE);
    }
}


/**
 * Samples for ImportPipelines Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ImportPipelineDelete.json
     */
    /**
     * Sample code: ImportPipelineDelete.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        importPipelineDelete(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getImportPipelines().delete("myResourceGroup", "myRegistry", "myImportPipeline",
            com.azure.core.util.Context.NONE);
    }
}

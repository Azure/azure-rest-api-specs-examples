
/**
 * Samples for ExportPipelines Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ExportPipelineDelete.json
     */
    /**
     * Sample code: ExportPipelineDelete.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        exportPipelineDelete(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getExportPipelines().delete("myResourceGroup", "myRegistry", "myExportPipeline",
            com.azure.core.util.Context.NONE);
    }
}

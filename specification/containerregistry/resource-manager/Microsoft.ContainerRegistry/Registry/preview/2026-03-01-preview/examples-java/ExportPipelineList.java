
/**
 * Samples for ExportPipelines List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ExportPipelineList.json
     */
    /**
     * Sample code: ExportPipelineList.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        exportPipelineList(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getExportPipelines().list("myResourceGroup", "myRegistry",
            com.azure.core.util.Context.NONE);
    }
}

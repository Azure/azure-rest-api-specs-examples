
/**
 * Samples for ExportPipelines Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ExportPipelineGet.json
     */
    /**
     * Sample code: ExportPipelineGet.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void exportPipelineGet(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getExportPipelines().getWithResponse("myResourceGroup", "myRegistry",
            "myExportPipeline", com.azure.core.util.Context.NONE);
    }
}

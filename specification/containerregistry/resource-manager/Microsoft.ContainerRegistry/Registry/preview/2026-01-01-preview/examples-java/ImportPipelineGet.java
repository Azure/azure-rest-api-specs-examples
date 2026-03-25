
/**
 * Samples for ImportPipelines Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ImportPipelineGet.json
     */
    /**
     * Sample code: ImportPipelineGet.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void importPipelineGet(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getImportPipelines().getWithResponse("myResourceGroup", "myRegistry",
            "myImportPipeline", com.azure.core.util.Context.NONE);
    }
}

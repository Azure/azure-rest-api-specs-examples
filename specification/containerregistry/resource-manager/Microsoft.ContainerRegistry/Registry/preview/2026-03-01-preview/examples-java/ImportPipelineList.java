
/**
 * Samples for ImportPipelines List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ImportPipelineList.json
     */
    /**
     * Sample code: ImportPipelineList.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        importPipelineList(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getImportPipelines().list("myResourceGroup", "myRegistry",
            com.azure.core.util.Context.NONE);
    }
}

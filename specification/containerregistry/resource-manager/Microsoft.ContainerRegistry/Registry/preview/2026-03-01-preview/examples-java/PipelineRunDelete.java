
/**
 * Samples for PipelineRuns Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/PipelineRunDelete.json
     */
    /**
     * Sample code: PipelineRunDelete.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void pipelineRunDelete(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getPipelineRuns().delete("myResourceGroup", "myRegistry", "myPipelineRun",
            com.azure.core.util.Context.NONE);
    }
}

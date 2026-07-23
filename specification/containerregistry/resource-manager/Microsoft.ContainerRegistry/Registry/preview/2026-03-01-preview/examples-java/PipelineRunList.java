
/**
 * Samples for PipelineRuns List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/PipelineRunList.json
     */
    /**
     * Sample code: PipelineRunList.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void pipelineRunList(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getPipelineRuns().list("myResourceGroup", "myRegistry",
            com.azure.core.util.Context.NONE);
    }
}

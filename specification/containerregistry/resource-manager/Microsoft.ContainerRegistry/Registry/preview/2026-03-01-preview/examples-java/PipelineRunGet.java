
/**
 * Samples for PipelineRuns Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/PipelineRunGet.json
     */
    /**
     * Sample code: PipelineRunGet.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void pipelineRunGet(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getPipelineRuns().getWithResponse("myResourceGroup", "myRegistry", "myPipelineRun",
            com.azure.core.util.Context.NONE);
    }
}

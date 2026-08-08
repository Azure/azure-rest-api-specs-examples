
/**
 * Samples for ModelSources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/ModelSources_Get.json
     */
    /**
     * Sample code: ModelSources_Get_MaximumSet.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void modelSourcesGetMaximumSet(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.modelSources().getWithResponse("rgaimanagers", "aimanager1", "huggingface",
            com.azure.core.util.Context.NONE);
    }
}

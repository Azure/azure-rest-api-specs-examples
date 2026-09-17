
/**
 * Samples for CustomAIModels Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-02-preview/CustomAIModels_Delete.json
     */
    /**
     * Sample code: Delete a CustomAIModel.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void deleteACustomAIModel(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.customAIModels().delete("rg1", "aimanager1", "custom-model1", null, com.azure.core.util.Context.NONE);
    }
}

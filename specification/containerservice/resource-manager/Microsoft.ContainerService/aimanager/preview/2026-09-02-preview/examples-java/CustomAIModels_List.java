
/**
 * Samples for CustomAIModels List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-02-preview/CustomAIModels_List.json
     */
    /**
     * Sample code: List CustomAIModel resources by AIManager.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void listCustomAIModelResourcesByAIManager(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.customAIModels().list("rg1", "aimanager1", com.azure.core.util.Context.NONE);
    }
}

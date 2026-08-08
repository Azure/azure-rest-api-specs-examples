
/**
 * Samples for ModelSources Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/ModelSources_Delete.json
     */
    /**
     * Sample code: ModelSources_Delete_MaximumSet.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void modelSourcesDeleteMaximumSet(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.modelSources().delete("rgaimanagers", "aimanager1", "huggingface", "\"abc123def456\"",
            com.azure.core.util.Context.NONE);
    }
}

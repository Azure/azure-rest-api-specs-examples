
/**
 * Samples for ModelSources List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/ModelSources_List.json
     */
    /**
     * Sample code: ModelSources_List_MaximumSet.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void modelSourcesListMaximumSet(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.modelSources().list("rgaimanagers", "aimanager1", com.azure.core.util.Context.NONE);
    }
}

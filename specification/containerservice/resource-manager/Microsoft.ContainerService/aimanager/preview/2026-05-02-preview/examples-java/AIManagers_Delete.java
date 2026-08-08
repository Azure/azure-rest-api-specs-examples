
/**
 * Samples for AIManagers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/AIManagers_Delete.json
     */
    /**
     * Sample code: Deletes an AI Manager resource.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void deletesAnAIManagerResource(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.aIManagers().delete("rg1", "aimanager1", null, com.azure.core.util.Context.NONE);
    }
}

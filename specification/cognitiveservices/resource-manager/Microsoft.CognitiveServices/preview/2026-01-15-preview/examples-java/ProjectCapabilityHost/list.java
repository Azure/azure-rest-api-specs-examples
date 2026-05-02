
/**
 * Samples for ProjectCapabilityHosts List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/ProjectCapabilityHost/list.json
     */
    /**
     * Sample code: List Project CapabilityHosts.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listProjectCapabilityHosts(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.projectCapabilityHosts().list("test-rg", "account-1", "project-1", com.azure.core.util.Context.NONE);
    }
}

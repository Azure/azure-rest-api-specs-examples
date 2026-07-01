
/**
 * Samples for ProjectCapabilityHosts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/ProjectCapabilityHost/delete.json
     */
    /**
     * Sample code: Delete Project CapabilityHost.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        deleteProjectCapabilityHost(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.projectCapabilityHosts().delete("test-rg", "account-1", "project-1", "capabilityHostName",
            com.azure.core.util.Context.NONE);
    }
}

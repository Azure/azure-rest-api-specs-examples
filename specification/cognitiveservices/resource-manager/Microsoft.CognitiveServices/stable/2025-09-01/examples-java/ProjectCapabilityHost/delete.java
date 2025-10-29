
/**
 * Samples for ProjectCapabilityHosts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * ProjectCapabilityHost/delete.json
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

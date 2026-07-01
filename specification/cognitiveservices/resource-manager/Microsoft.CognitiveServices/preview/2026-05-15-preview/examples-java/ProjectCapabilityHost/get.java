
/**
 * Samples for ProjectCapabilityHosts Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/ProjectCapabilityHost/get.json
     */
    /**
     * Sample code: Get Project CapabilityHost.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getProjectCapabilityHost(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.projectCapabilityHosts().getWithResponse("test-rg", "account-1", "project-1", "capabilityHostName",
            com.azure.core.util.Context.NONE);
    }
}

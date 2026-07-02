
/**
 * Samples for AgentApplications Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/AgentApplication/get.json
     */
    /**
     * Sample code: Get Account Agent Application.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getAccountAgentApplication(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.agentApplications().getWithResponse("test-rg", "my-cognitive-services-account", "my-project",
            "agent-app-1", com.azure.core.util.Context.NONE);
    }
}

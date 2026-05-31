
/**
 * Samples for AgentApplications Enable.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/AgentApplication/enable.json
     */
    /**
     * Sample code: Enable Agent Application.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        enableAgentApplication(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.agentApplications().enableWithResponse("test-rg", "my-cognitive-services-account", "my-project",
            "agent-app-1", com.azure.core.util.Context.NONE);
    }
}

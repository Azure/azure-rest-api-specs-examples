
/**
 * Samples for AgentDeployments List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/AgentDeployment/list.json
     */
    /**
     * Sample code: List Agent Deployments.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listAgentDeployments(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.agentDeployments().list("test-rg", "my-cognitive-services-account", "my-project", "agent-app-1", null,
            null, null, null, null, com.azure.core.util.Context.NONE);
    }
}

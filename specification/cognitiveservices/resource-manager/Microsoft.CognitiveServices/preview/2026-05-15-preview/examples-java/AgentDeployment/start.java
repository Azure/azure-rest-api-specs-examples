
/**
 * Samples for AgentDeployments Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/AgentDeployment/start.json
     */
    /**
     * Sample code: Start Agent Deployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        startAgentDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.agentDeployments().startWithResponse("test-rg", "my-cognitive-services-account", "my-project",
            "agent-app-1", "deployment-1", com.azure.core.util.Context.NONE);
    }
}

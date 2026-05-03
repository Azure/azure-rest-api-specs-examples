
/**
 * Samples for AgentDeployments Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/AgentDeployment/stop.json
     */
    /**
     * Sample code: Stop Agent Deployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        stopAgentDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.agentDeployments().stopWithResponse("test-rg", "my-cognitive-services-account", "my-project",
            "agent-app-1", "deployment-1", com.azure.core.util.Context.NONE);
    }
}

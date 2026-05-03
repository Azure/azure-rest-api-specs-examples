
/**
 * Samples for AgentDeployments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/AgentDeployment/delete.json
     */
    /**
     * Sample code: Delete Agent Deployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        deleteAgentDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.agentDeployments().delete("test-rg", "my-cognitive-services-account", "my-project", "agent-app-1",
            "deployment-1", com.azure.core.util.Context.NONE);
    }
}

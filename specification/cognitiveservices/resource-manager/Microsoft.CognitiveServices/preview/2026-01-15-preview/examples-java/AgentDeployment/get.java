
/**
 * Samples for AgentDeployments Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/AgentDeployment/get.json
     */
    /**
     * Sample code: Get Agent Deployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getAgentDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.agentDeployments().getWithResponse("test-rg", "my-cognitive-services-account", "my-project",
            "agent-app-1", "deployment-1", com.azure.core.util.Context.NONE);
    }
}

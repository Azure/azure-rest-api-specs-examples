
/**
 * Samples for AgentApplications ListAgents.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/AgentApplication/listAgents.json
     */
    /**
     * Sample code: List Agents for Agent Application.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listAgentsForAgentApplication(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.agentApplications().listAgentsWithResponse("test-rg", "my-cognitive-services-account", "my-project",
            "agent-app-1", com.azure.core.util.Context.NONE);
    }
}

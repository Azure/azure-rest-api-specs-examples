
/**
 * Samples for AgentApplications Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/AgentApplication/delete.json
     */
    /**
     * Sample code: Delete Account Agent Application.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        deleteAccountAgentApplication(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.agentApplications().delete("test-rg", "my-cognitive-services-account", "my-project", "agent-app-1",
            com.azure.core.util.Context.NONE);
    }
}

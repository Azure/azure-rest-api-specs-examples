
import java.util.Arrays;

/**
 * Samples for AgentApplications List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/AgentApplication/list.json
     */
    /**
     * Sample code: List Account Agent Applications.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listAccountAgentApplications(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.agentApplications().list("test-rg", "my-cognitive-services-account", "my-project", 30, null, "string",
            Arrays.asList("agent-app-1", "agent-app-2"), "test", "name", true, com.azure.core.util.Context.NONE);
    }
}

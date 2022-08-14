import com.azure.core.util.Context;

/** Samples for AgentRegistrationInformation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getAgentRegistration.json
     */
    /**
     * Sample code: Get the agent registration information.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getTheAgentRegistrationInformation(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.agentRegistrationInformations().getWithResponse("rg", "myAutomationAccount18", Context.NONE);
    }
}

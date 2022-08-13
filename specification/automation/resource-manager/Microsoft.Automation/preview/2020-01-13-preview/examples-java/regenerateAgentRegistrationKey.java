import com.azure.core.util.Context;
import com.azure.resourcemanager.automation.models.AgentRegistrationKeyName;
import com.azure.resourcemanager.automation.models.AgentRegistrationRegenerateKeyParameter;

/** Samples for AgentRegistrationInformation RegenerateKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/regenerateAgentRegistrationKey.json
     */
    /**
     * Sample code: Regenerate registration key.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void regenerateRegistrationKey(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .agentRegistrationInformations()
            .regenerateKeyWithResponse(
                "rg",
                "myAutomationAccount18",
                new AgentRegistrationRegenerateKeyParameter().withKeyName(AgentRegistrationKeyName.PRIMARY),
                Context.NONE);
    }
}

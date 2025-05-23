
import com.azure.resourcemanager.recoveryservicessiterecovery.models.A2AUpdateContainerMappingInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.AgentAutoUpdateStatus;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.ProtectionContainerMapping;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.UpdateProtectionContainerMappingInputProperties;

/**
 * Samples for ReplicationProtectionContainerMappings Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationProtectionContainerMappings_Update.json
     */
    /**
     * Sample code: Update protection container mapping.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void updateProtectionContainerMapping(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        ProtectionContainerMapping resource = manager.replicationProtectionContainerMappings()
            .getWithResponse("resourceGroupPS1", "vault1", "cloud1", "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
                "cloud1protectionprofile1", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withProperties(new UpdateProtectionContainerMappingInputProperties()
            .withProviderSpecificInput(new A2AUpdateContainerMappingInput()
                .withAgentAutoUpdateStatus(AgentAutoUpdateStatus.ENABLED).withAutomationAccountArmId(
                    "/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/automationrg1/providers/Microsoft.Automation/automationAccounts/automationaccount1")))
            .apply();
    }
}

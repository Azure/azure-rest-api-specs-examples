
import com.azure.resourcemanager.recoveryservicessiterecovery.models.CreateProtectionContainerMappingInputProperties;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.ReplicationProviderSpecificContainerMappingInput;

/**
 * Samples for ReplicationProtectionContainerMappings Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationProtectionContainerMappings_Create.json
     */
    /**
     * Sample code: Create protection container mapping.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void createProtectionContainerMapping(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectionContainerMappings().define("cloud1protectionprofile1")
            .withExistingReplicationProtectionContainer("resourceGroupPS1", "vault1", "cloud1",
                "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179")
            .withProperties(new CreateProtectionContainerMappingInputProperties()
                .withTargetProtectionContainerId("Microsoft Azure")
                .withPolicyId(
                    "/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/protectionprofile1")
                .withProviderSpecificInput(new ReplicationProviderSpecificContainerMappingInput()))
            .create();
    }
}


import com.azure.resourcemanager.recoveryservicessiterecovery.models.HyperVReplicaAzurePolicyInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.Policy;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.UpdatePolicyInputProperties;

/**
 * Samples for ReplicationPolicies Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples
     * /ReplicationPolicies_Update.json
     */
    /**
     * Sample code: Updates the policy.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        updatesThePolicy(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        Policy resource = manager.replicationPolicies()
            .getWithResponse("vault1", "resourceGroupPS1", "protectionprofile1", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update()
            .withProperties(
                new UpdatePolicyInputProperties().withReplicationProviderSettings(new HyperVReplicaAzurePolicyInput()))
            .apply();
    }
}

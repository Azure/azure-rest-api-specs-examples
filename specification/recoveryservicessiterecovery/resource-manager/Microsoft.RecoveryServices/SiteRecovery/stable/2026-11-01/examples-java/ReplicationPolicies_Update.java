
import com.azure.resourcemanager.recoveryservicessiterecovery.models.HyperVReplicaAzurePolicyInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.Policy;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.UpdatePolicyInputProperties;

/**
 * Samples for ReplicationPolicies Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationPolicies_Update.json
     */
    /**
     * Sample code: Updates the policy.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        updatesThePolicy(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        Policy resource = manager.replicationPolicies()
            .getWithResponse("resourceGroupPS1", "vault1", "protectionprofile1", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update()
            .withProperties(
                new UpdatePolicyInputProperties().withReplicationProviderSettings(new HyperVReplicaAzurePolicyInput()))
            .apply();
    }
}

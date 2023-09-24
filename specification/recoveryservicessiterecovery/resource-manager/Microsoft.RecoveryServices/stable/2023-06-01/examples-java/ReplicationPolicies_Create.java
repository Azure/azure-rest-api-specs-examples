import com.azure.resourcemanager.recoveryservicessiterecovery.models.CreatePolicyInputProperties;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.HyperVReplicaAzurePolicyInput;

/** Samples for ReplicationPolicies Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationPolicies_Create.json
     */
    /**
     * Sample code: Creates the policy.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void createsThePolicy(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .replicationPolicies()
            .define("protectionprofile1")
            .withExistingVault("vault1", "resourceGroupPS1")
            .withProperties(
                new CreatePolicyInputProperties().withProviderSpecificInput(new HyperVReplicaAzurePolicyInput()))
            .create();
    }
}

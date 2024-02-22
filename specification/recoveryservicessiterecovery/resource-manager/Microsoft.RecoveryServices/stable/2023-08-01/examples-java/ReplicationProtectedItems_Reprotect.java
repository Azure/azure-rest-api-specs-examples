
import com.azure.resourcemanager.recoveryservicessiterecovery.models.HyperVReplicaAzureReprotectInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.ReverseReplicationInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.ReverseReplicationInputProperties;

/**
 * Samples for ReplicationProtectedItems Reprotect.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples
     * /ReplicationProtectedItems_Reprotect.json
     */
    /**
     * Sample code: Execute Reverse Replication\Reprotect.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void executeReverseReplicationReprotect(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectedItems().reprotect("vault1", "resourceGroupPS1", "cloud1",
            "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", "f8491e4f-817a-40dd-a90c-af773978c75b",
            new ReverseReplicationInput()
                .withProperties(new ReverseReplicationInputProperties().withFailoverDirection("PrimaryToRecovery")
                    .withProviderSpecificDetails(new HyperVReplicaAzureReprotectInput())),
            com.azure.core.util.Context.NONE);
    }
}


import com.azure.resourcemanager.recoveryservicessiterecovery.models.HyperVReplicaAzureUnplannedFailoverInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.UnplannedFailoverInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.UnplannedFailoverInputProperties;

/**
 * Samples for ReplicationProtectedItems UnplannedFailover.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationProtectedItems_UnplannedFailover.json
     */
    /**
     * Sample code: Execute unplanned failover.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        executeUnplannedFailover(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectedItems().unplannedFailover("resourceGroupPS1", "vault1", "cloud1",
            "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", "f8491e4f-817a-40dd-a90c-af773978c75b",
            new UnplannedFailoverInput().withProperties(new UnplannedFailoverInputProperties()
                .withFailoverDirection("PrimaryToRecovery").withSourceSiteOperations("NotRequired")
                .withProviderSpecificDetails(new HyperVReplicaAzureUnplannedFailoverInput())),
            com.azure.core.util.Context.NONE);
    }
}


import com.azure.resourcemanager.recoveryservicessiterecovery.models.HyperVReplicaAzureTestFailoverInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.TestFailoverInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.TestFailoverInputProperties;

/**
 * Samples for ReplicationProtectedItems TestFailover.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationProtectedItems_TestFailover.json
     */
    /**
     * Sample code: Execute test failover.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        executeTestFailover(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectedItems().testFailover("resourceGroupPS1", "vault1", "cloud1",
            "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", "f8491e4f-817a-40dd-a90c-af773978c75b",
            new TestFailoverInput().withProperties(new TestFailoverInputProperties()
                .withFailoverDirection("PrimaryToRecovery").withNetworkType("VmNetworkAsInput")
                .withNetworkId(
                    "/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/siterecoveryProd1/providers/Microsoft.Network/virtualNetworks/vnetavrai")
                .withProviderSpecificDetails(new HyperVReplicaAzureTestFailoverInput())),
            com.azure.core.util.Context.NONE);
    }
}

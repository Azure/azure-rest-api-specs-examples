
import com.azure.resourcemanager.recoveryservicessiterecovery.models.InMageAzureV2SwitchProviderInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.SwitchProviderInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.SwitchProviderInputProperties;

/**
 * Samples for ReplicationProtectedItems SwitchProvider.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationProtectedItems_SwitchProvider.json
     */
    /**
     * Sample code: Execute switch provider.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        executeSwitchProvider(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectedItems().switchProvider("resourceGroupPS1", "vault1", "cloud1",
            "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", "f8491e4f-817a-40dd-a90c-af773978c75b",
            new SwitchProviderInput().withProperties(new SwitchProviderInputProperties()
                .withTargetInstanceType("InMageRcm")
                .withProviderSpecificDetails(new InMageAzureV2SwitchProviderInput().withTargetVaultId(
                    "/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault2")
                    .withTargetFabricId(
                        "/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud2")
                    .withTargetApplianceId("5efaa202-e958-435e-8171-706bf735fcc4"))),
            com.azure.core.util.Context.NONE);
    }
}

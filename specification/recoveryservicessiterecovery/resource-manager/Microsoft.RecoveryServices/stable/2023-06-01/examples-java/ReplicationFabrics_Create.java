import com.azure.resourcemanager.recoveryservicessiterecovery.models.FabricCreationInputProperties;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.FabricSpecificCreationInput;

/** Samples for ReplicationFabrics Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationFabrics_Create.json
     */
    /**
     * Sample code: Creates an Azure Site Recovery fabric.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void createsAnAzureSiteRecoveryFabric(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .replicationFabrics()
            .define("cloud1")
            .withExistingVault("vault1", "resourceGroupPS1")
            .withProperties(new FabricCreationInputProperties().withCustomDetails(new FabricSpecificCreationInput()))
            .create();
    }
}

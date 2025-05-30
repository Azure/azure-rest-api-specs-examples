
import com.azure.resourcemanager.recoveryservicessiterecovery.models.VaultSettingCreationInputProperties;

/**
 * Samples for ReplicationVaultSetting Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationVaultSetting_Create.json
     */
    /**
     * Sample code: Updates vault setting. A vault setting object is a singleton per vault and it is always present by
     * default.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void updatesVaultSettingAVaultSettingObjectIsASingletonPerVaultAndItIsAlwaysPresentByDefault(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationVaultSettings().define("default").withExistingVault("resourceGroupPS1", "vault1")
            .withProperties(new VaultSettingCreationInputProperties().withMigrationSolutionId(
                "/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.Migrate/MigrateProjects/resourceGroupPS1-MigrateProject/Solutions/Servers-Migration-ServerMigration"))
            .create();
    }
}

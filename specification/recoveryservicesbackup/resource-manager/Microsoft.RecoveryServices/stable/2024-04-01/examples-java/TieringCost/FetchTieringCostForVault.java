
import com.azure.resourcemanager.recoveryservicesbackup.models.FetchTieringCostSavingsInfoForVaultRequest;
import com.azure.resourcemanager.recoveryservicesbackup.models.RecoveryPointTierType;

/**
 * Samples for FetchTieringCost Post.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/
     * TieringCost/FetchTieringCostForVault.json
     */
    /**
     * Sample code: Get the tiering savings cost info for vault.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getTheTieringSavingsCostInfoForVault(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.fetchTieringCosts().post("netsdktestrg", "testVault",
            new FetchTieringCostSavingsInfoForVaultRequest().withSourceTierType(RecoveryPointTierType.HARDENED_RP)
                .withTargetTierType(RecoveryPointTierType.ARCHIVED_RP),
            com.azure.core.util.Context.NONE);
    }
}

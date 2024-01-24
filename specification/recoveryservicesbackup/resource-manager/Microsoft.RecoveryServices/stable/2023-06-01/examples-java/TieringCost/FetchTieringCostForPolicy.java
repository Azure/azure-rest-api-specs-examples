
import com.azure.resourcemanager.recoveryservicesbackup.models.FetchTieringCostSavingsInfoForPolicyRequest;
import com.azure.resourcemanager.recoveryservicesbackup.models.RecoveryPointTierType;

/**
 * Samples for FetchTieringCost Post.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/
     * TieringCost/FetchTieringCostForPolicy.json
     */
    /**
     * Sample code: Get the tiering savings cost info for policy.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getTheTieringSavingsCostInfoForPolicy(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.fetchTieringCosts().post("netsdktestrg", "testVault",
            new FetchTieringCostSavingsInfoForPolicyRequest().withSourceTierType(RecoveryPointTierType.HARDENED_RP)
                .withTargetTierType(RecoveryPointTierType.ARCHIVED_RP).withPolicyName("monthly"),
            com.azure.core.util.Context.NONE);
    }
}

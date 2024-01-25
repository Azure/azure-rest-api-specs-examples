
import com.azure.resourcemanager.recoveryservicesbackup.models.FetchTieringCostInfoForRehydrationRequest;
import com.azure.resourcemanager.recoveryservicesbackup.models.RecoveryPointTierType;
import com.azure.resourcemanager.recoveryservicesbackup.models.RehydrationPriority;

/**
 * Samples for FetchTieringCost Post.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/
     * TieringCost/FetchTieringCostForRehydrate.json
     */
    /**
     * Sample code: Get the rehydration cost for recovery point.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getTheRehydrationCostForRecoveryPoint(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.fetchTieringCosts().post("netsdktestrg", "testVault",
            new FetchTieringCostInfoForRehydrationRequest().withSourceTierType(RecoveryPointTierType.ARCHIVED_RP)
                .withTargetTierType(RecoveryPointTierType.HARDENED_RP)
                .withContainerName("IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1")
                .withProtectedItemName("VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1")
                .withRecoveryPointId("1222343434").withRehydrationPriority(RehydrationPriority.HIGH),
            com.azure.core.util.Context.NONE);
    }
}

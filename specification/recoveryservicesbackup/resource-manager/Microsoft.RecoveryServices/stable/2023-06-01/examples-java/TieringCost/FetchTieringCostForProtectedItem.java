
import com.azure.resourcemanager.recoveryservicesbackup.models.FetchTieringCostSavingsInfoForProtectedItemRequest;
import com.azure.resourcemanager.recoveryservicesbackup.models.RecoveryPointTierType;

/**
 * Samples for FetchTieringCost Post.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/
     * TieringCost/FetchTieringCostForProtectedItem.json
     */
    /**
     * Sample code: Get the tiering savings cost info for protected item.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getTheTieringSavingsCostInfoForProtectedItem(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.fetchTieringCosts().post("netsdktestrg", "testVault",
            new FetchTieringCostSavingsInfoForProtectedItemRequest()
                .withSourceTierType(RecoveryPointTierType.HARDENED_RP)
                .withTargetTierType(RecoveryPointTierType.ARCHIVED_RP)
                .withContainerName("IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1")
                .withProtectedItemName("VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1"),
            com.azure.core.util.Context.NONE);
    }
}

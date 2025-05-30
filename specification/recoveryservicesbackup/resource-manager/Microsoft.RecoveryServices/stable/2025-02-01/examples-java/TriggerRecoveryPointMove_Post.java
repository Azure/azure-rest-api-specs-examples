
import com.azure.resourcemanager.recoveryservicesbackup.models.MoveRPAcrossTiersRequest;
import com.azure.resourcemanager.recoveryservicesbackup.models.RecoveryPointTierType;

/**
 * Samples for ResourceProvider MoveRecoveryPoint.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/
     * TriggerRecoveryPointMove_Post.json
     */
    /**
     * Sample code: Trigger RP Move Operation.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void
        triggerRPMoveOperation(com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.resourceProviders().moveRecoveryPoint("testVault", "netsdktestrg", "Azure",
            "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
            "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1", "348916168024334",
            new MoveRPAcrossTiersRequest().withObjectType("MoveRPAcrossTiersRequest")
                .withSourceTierType(RecoveryPointTierType.HARDENED_RP)
                .withTargetTierType(RecoveryPointTierType.ARCHIVED_RP),
            com.azure.core.util.Context.NONE);
    }
}

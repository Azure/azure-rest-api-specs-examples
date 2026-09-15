
import com.azure.resourcemanager.recoveryservicesbackup.models.InstantItemRecoveryOperationResultRequest;

/**
 * Samples for ItemLevelRecoveryConnections ListInstantItemRecoveryOperationResult.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01/AzureIaasVm/ListInstantItemRecoveryOperationResult.json
     */
    /**
     * Sample code: List the Instant Item Recovery operation result (mount scripts) for an active ILR session.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listTheInstantItemRecoveryOperationResultMountScriptsForAnActiveILRSession(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.itemLevelRecoveryConnections().listInstantItemRecoveryOperationResultWithResponse(
            "PythonSDKBackupTestRg", "PySDKBackupTestRsVault", "Azure",
            "iaasvmcontainer;iaasvmcontainerv2;pysdktestrg;pysdktestv2vm1",
            "vm;iaasvmcontainerv2;pysdktestrg;pysdktestv2vm1", "38823086363464",
            new InstantItemRecoveryOperationResultRequest().withProvisionInstantItemRecoveryOperationId(
                "00000000-0000-0000-0000-000000000001"),
            com.azure.core.util.Context.NONE);
    }
}

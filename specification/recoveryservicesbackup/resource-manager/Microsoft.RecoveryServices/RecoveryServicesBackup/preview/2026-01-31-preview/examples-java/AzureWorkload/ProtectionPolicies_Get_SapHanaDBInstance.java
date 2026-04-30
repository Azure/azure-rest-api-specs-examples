
/**
 * Samples for ProtectionPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/AzureWorkload/ProtectionPolicies_Get_SapHanaDBInstance.json
     */
    /**
     * Sample code: Get Sap Hana DB Instance Workload Protection Policy Details.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getSapHanaDBInstanceWorkloadProtectionPolicyDetails(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.protectionPolicies().getWithResponse("HanaTestRsVault", "SwaggerTestRg", "testHanaSnapshotV2Policy1",
            com.azure.core.util.Context.NONE);
    }
}

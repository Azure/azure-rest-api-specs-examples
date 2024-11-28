
/**
 * Samples for CloudHsmClusterBackupStatus Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2024-06-30-
     * preview/examples/CloudHsmCluster_Backup_Pending_MaximumSet_Gen.json
     */
    /**
     * Sample code: CloudHsmCluster_Get_Backup_Status_MaximumSet_Gen.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void cloudHsmClusterGetBackupStatusMaximumSetGen(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.cloudHsmClusterBackupStatus().getWithResponse("rgcloudhsm", "chsm1", "572a45927fc240e1ac075de27371680b",
            com.azure.core.util.Context.NONE);
    }
}

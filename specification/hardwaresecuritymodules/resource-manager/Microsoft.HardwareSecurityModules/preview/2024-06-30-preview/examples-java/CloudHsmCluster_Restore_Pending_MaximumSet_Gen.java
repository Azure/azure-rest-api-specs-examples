
/**
 * Samples for CloudHsmClusterRestoreStatus Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2024-06-30-
     * preview/examples/CloudHsmCluster_Restore_Pending_MaximumSet_Gen.json
     */
    /**
     * Sample code: CloudHsmCluster_Get_Restore_Status_MaximumSet_Gen.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void cloudHsmClusterGetRestoreStatusMaximumSetGen(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.cloudHsmClusterRestoreStatus().getWithResponse("rgcloudhsm", "chsm1",
            "572a45927fc240e1ac075de27371680b", com.azure.core.util.Context.NONE);
    }
}

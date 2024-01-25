
/**
 * Samples for CloudHsmClusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2023-12-10-
     * preview/examples/CloudHsmCluster_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: CloudHsmCluster_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void cloudHsmClusterDeleteMaximumSetGen(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.cloudHsmClusters().delete("rgcloudhsm", "chsm1", com.azure.core.util.Context.NONE);
    }
}

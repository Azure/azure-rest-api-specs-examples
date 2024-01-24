
/**
 * Samples for CloudHsmClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2023-12-10-
     * preview/examples/CloudHsmCluster_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: CloudHsmCluster_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void cloudHsmClusterListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.cloudHsmClusters().list(null, com.azure.core.util.Context.NONE);
    }
}

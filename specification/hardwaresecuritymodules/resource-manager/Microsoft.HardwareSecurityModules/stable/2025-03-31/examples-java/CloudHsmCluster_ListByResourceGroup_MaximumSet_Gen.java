
/**
 * Samples for CloudHsmClusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-31/CloudHsmCluster_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: CloudHsmCluster_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void cloudHsmClusterListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.cloudHsmClusters().listByResourceGroup("rgcloudhsm", null, com.azure.core.util.Context.NONE);
    }
}

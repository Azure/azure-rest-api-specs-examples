
/**
 * Samples for CloudHsmClusterPrivateLinkResources ListByCloudHsmCluster.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2024-06-30-
     * preview/examples/CloudHsmClusterPrivateLinkResource_ListByCloudHsmCluster_MaximumSet_Gen.json
     */
    /**
     * Sample code: CloudHsmClusterPrivateLinkResources_ListByResource_MaximumSet_Gen.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void cloudHsmClusterPrivateLinkResourcesListByResourceMaximumSetGen(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.cloudHsmClusterPrivateLinkResources().listByCloudHsmCluster("rgcloudhsm", "chsm1",
            com.azure.core.util.Context.NONE);
    }
}

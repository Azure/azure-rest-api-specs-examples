
/**
 * Samples for SapVirtualInstances GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/
     * examples/sapvirtualinstances/SAPVirtualInstances_Get_ACSSInstallationBlocked.json
     */
    /**
     * Sample code: SAPVirtualInstances Get With ACSS Installation Blocked.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sAPVirtualInstancesGetWithACSSInstallationBlocked(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapVirtualInstances().getByResourceGroupWithResponse("test-rg", "X00",
            com.azure.core.util.Context.NONE);
    }
}

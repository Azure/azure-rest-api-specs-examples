
/**
 * Samples for SapVirtualInstances GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapVirtualInstances_GetAcssInstallationBlocked.json
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


/**
 * Samples for SapVirtualInstances GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapVirtualInstances_Get.json
     */
    /**
     * Sample code: SAPVirtualInstances_Get.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sAPVirtualInstancesGet(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapVirtualInstances().getByResourceGroupWithResponse("test-rg", "X00",
            com.azure.core.util.Context.NONE);
    }
}

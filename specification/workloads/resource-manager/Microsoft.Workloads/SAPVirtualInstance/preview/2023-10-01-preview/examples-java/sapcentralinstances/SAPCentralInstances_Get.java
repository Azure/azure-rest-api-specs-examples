
/**
 * Samples for SapCentralInstances Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/
     * examples/sapcentralinstances/SAPCentralInstances_Get.json
     */
    /**
     * Sample code: SAPCentralInstances_Get.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sAPCentralInstancesGet(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapCentralInstances().getWithResponse("test-rg", "X00", "centralServer",
            com.azure.core.util.Context.NONE);
    }
}

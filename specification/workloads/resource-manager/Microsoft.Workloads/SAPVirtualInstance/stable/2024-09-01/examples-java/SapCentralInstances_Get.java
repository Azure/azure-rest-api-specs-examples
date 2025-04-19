
/**
 * Samples for SapCentralServerInstances Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapCentralInstances_Get.json
     */
    /**
     * Sample code: SapCentralServerInstances_Get.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sapCentralServerInstancesGet(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapCentralServerInstances().getWithResponse("test-rg", "X00", "centralServer",
            com.azure.core.util.Context.NONE);
    }
}

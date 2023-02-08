/** Samples for SapCentralInstances Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2022-11-01-preview/examples/sapvirtualinstances/SAPCentralInstances_Get.json
     */
    /**
     * Sample code: SAPCentralInstances_Get.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPCentralInstancesGet(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .sapCentralInstances()
            .getWithResponse("test-rg", "X00", "centralServer", com.azure.core.util.Context.NONE);
    }
}

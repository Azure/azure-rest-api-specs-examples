/** Samples for SapDatabaseInstances Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2022-11-01-preview/examples/sapvirtualinstances/SAPDatabaseInstances_Get.json
     */
    /**
     * Sample code: SAPDatabaseInstances_Get.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPDatabaseInstancesGet(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .sapDatabaseInstances()
            .getWithResponse("test-rg", "X00", "databaseServer", com.azure.core.util.Context.NONE);
    }
}

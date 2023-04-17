/** Samples for SapDatabaseInstances Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPDatabaseInstances_Get.json
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

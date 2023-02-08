/** Samples for SapDatabaseInstances Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2022-11-01-preview/examples/sapvirtualinstances/SAPDatabaseInstances_Delete.json
     */
    /**
     * Sample code: SAPDatabaseInstances_Delete.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPDatabaseInstancesDelete(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.sapDatabaseInstances().delete("test-rg", "X00", "databaseServer", com.azure.core.util.Context.NONE);
    }
}

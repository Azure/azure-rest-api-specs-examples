import com.azure.core.util.Context;

/** Samples for SapDatabaseInstances List. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPDatabaseInstances_List.json
     */
    /**
     * Sample code: SAPDatabaseInstances_List.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPDatabaseInstancesList(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.sapDatabaseInstances().list("test-rg", "X00", Context.NONE);
    }
}

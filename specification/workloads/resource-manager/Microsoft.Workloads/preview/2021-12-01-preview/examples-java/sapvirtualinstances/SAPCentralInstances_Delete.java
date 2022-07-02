import com.azure.core.util.Context;

/** Samples for SapCentralInstances Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPCentralInstances_Delete.json
     */
    /**
     * Sample code: SAPCentralInstances_Delete.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPCentralInstancesDelete(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.sapCentralInstances().delete("test-rg", "X00", "centralServer", Context.NONE);
    }
}

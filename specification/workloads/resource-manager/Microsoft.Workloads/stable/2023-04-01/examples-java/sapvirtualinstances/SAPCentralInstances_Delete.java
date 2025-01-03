
/**
 * Samples for SapCentralInstances Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/
     * SAPCentralInstances_Delete.json
     */
    /**
     * Sample code: SAPCentralInstances_Delete.
     * 
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPCentralInstancesDelete(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.sapCentralInstances().delete("test-rg", "X00", "centralServer", com.azure.core.util.Context.NONE);
    }
}

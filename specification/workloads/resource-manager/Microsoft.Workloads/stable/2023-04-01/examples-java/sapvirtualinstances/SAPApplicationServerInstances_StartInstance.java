
/**
 * Samples for SapApplicationServerInstances StartInstance.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/
     * SAPApplicationServerInstances_StartInstance.json
     */
    /**
     * Sample code: Start the SAP Application Server Instance.
     * 
     * @param manager Entry point to WorkloadsManager.
     */
    public static void
        startTheSAPApplicationServerInstance(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.sapApplicationServerInstances().startInstance("test-rg", "X00", "app01",
            com.azure.core.util.Context.NONE);
    }
}

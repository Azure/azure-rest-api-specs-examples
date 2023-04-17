/** Samples for SapApplicationServerInstances List. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPApplicationServerInstances_List.json
     */
    /**
     * Sample code: SAPApplicationServerInstances_List.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPApplicationServerInstancesList(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.sapApplicationServerInstances().list("test-rg", "X00", com.azure.core.util.Context.NONE);
    }
}

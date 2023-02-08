/** Samples for SapApplicationServerInstances List. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2022-11-01-preview/examples/sapvirtualinstances/SAPApplicationServerInstances_List.json
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

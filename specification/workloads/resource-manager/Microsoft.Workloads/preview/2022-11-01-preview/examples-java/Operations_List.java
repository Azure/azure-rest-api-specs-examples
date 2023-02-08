/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2022-11-01-preview/examples/Operations_List.json
     */
    /**
     * Sample code: Operations.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void operations(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}

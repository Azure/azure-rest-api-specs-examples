/** Samples for Labels Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/riskiq/resource-manager/Microsoft.Easm/preview/2023-04-01-preview/examples/Labels_Delete.json
     */
    /**
     * Sample code: Labels.
     *
     * @param manager Entry point to EasmManager.
     */
    public static void labels(com.azure.resourcemanager.defendereasm.EasmManager manager) {
        manager.labels().delete("dummyrg", "ThisisaWorkspace", "ThisisaLabel", com.azure.core.util.Context.NONE);
    }
}

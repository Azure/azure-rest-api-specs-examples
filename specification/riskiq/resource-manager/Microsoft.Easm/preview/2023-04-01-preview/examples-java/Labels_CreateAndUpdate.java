/** Samples for Labels CreateAndUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/riskiq/resource-manager/Microsoft.Easm/preview/2023-04-01-preview/examples/Labels_CreateAndUpdate.json
     */
    /**
     * Sample code: Labels.
     *
     * @param manager Entry point to EasmManager.
     */
    public static void labels(com.azure.resourcemanager.defendereasm.EasmManager manager) {
        manager
            .labels()
            .createAndUpdate("dummyrg", "ThisisaWorkspace", "ThisisaLabel", null, com.azure.core.util.Context.NONE);
    }
}

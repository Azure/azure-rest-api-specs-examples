/** Samples for Labels GetByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/riskiq/resource-manager/Microsoft.Easm/preview/2023-04-01-preview/examples/Labels_GetByWorkspace.json
     */
    /**
     * Sample code: Labels.
     *
     * @param manager Entry point to EasmManager.
     */
    public static void labels(com.azure.resourcemanager.defendereasm.EasmManager manager) {
        manager
            .labels()
            .getByWorkspaceWithResponse(
                "dummyrg", "ThisisaWorkspace", "ThisisaLabel", com.azure.core.util.Context.NONE);
    }
}

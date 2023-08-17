/** Samples for Tasks GetByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/riskiq/resource-manager/Microsoft.Easm/preview/2023-04-01-preview/examples/Tasks_GetByWorkspace.json
     */
    /**
     * Sample code: Tasks.
     *
     * @param manager Entry point to EasmManager.
     */
    public static void tasks(com.azure.resourcemanager.defendereasm.EasmManager manager) {
        manager
            .tasks()
            .getByWorkspaceWithResponse(
                "dummyrg", "ThisisaWorkspace", "ThisisaTaskId", com.azure.core.util.Context.NONE);
    }
}

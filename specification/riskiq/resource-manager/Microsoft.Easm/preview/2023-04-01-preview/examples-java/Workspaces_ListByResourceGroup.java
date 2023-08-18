/** Samples for Workspaces ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/riskiq/resource-manager/Microsoft.Easm/preview/2023-04-01-preview/examples/Workspaces_ListByResourceGroup.json
     */
    /**
     * Sample code: Workspaces.
     *
     * @param manager Entry point to EasmManager.
     */
    public static void workspaces(com.azure.resourcemanager.defendereasm.EasmManager manager) {
        manager.workspaces().listByResourceGroup("dummyrg", com.azure.core.util.Context.NONE);
    }
}

import com.azure.resourcemanager.defendereasm.models.WorkspaceResource;

/** Samples for Workspaces Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/riskiq/resource-manager/Microsoft.Easm/preview/2023-04-01-preview/examples/Workspaces_Update.json
     */
    /**
     * Sample code: Workspaces.
     *
     * @param manager Entry point to EasmManager.
     */
    public static void workspaces(com.azure.resourcemanager.defendereasm.EasmManager manager) {
        WorkspaceResource resource =
            manager
                .workspaces()
                .getByResourceGroupWithResponse("dummyrg", "ThisisaWorkspace", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().apply();
    }
}

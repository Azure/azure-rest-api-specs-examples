
/**
 * Samples for PermissionBindings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/PermissionBindings_Delete
     * .json
     */
    /**
     * Sample code: PermissionBindings_Delete.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void permissionBindingsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.permissionBindings().delete("examplerg", "exampleNamespaceName1", "examplePermissionBindingName1",
            com.azure.core.util.Context.NONE);
    }
}

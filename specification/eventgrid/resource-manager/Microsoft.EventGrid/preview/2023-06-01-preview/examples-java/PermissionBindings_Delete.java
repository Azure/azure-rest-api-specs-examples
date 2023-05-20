/** Samples for PermissionBindings Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/PermissionBindings_Delete.json
     */
    /**
     * Sample code: PermissionBindings_Delete.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void permissionBindingsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .permissionBindings()
            .delete(
                "examplerg",
                "exampleNamespaceName1",
                "examplePermissionBindingName1",
                com.azure.core.util.Context.NONE);
    }
}

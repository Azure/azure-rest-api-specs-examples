/** Samples for PermissionBindings ListByNamespace. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/PermissionBindings_ListByNamespace.json
     */
    /**
     * Sample code: PermissionBindings_ListByNamespace.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void permissionBindingsListByNamespace(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .permissionBindings()
            .listByNamespace("examplerg", "namespace123", null, null, com.azure.core.util.Context.NONE);
    }
}

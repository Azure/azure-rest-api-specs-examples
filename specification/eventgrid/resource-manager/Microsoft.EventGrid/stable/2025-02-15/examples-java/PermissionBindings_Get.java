
/**
 * Samples for PermissionBindings Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/PermissionBindings_Get.
     * json
     */
    /**
     * Sample code: PermissionBindings_Get.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void permissionBindingsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.permissionBindings().getWithResponse("examplerg", "exampleNamespaceName1",
            "examplePermissionBindingName1", com.azure.core.util.Context.NONE);
    }
}

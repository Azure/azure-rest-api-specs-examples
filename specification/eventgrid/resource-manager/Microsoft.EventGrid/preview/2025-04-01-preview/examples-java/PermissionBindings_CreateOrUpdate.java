
import com.azure.resourcemanager.eventgrid.models.PermissionType;

/**
 * Samples for PermissionBindings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/
     * PermissionBindings_CreateOrUpdate.json
     */
    /**
     * Sample code: PermissionBindings_CreateOrUpdate.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void permissionBindingsCreateOrUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.permissionBindings().define("examplePermissionBindingName1")
            .withExistingNamespace("examplerg", "exampleNamespaceName1").withTopicSpaceName("exampleTopicSpaceName1")
            .withPermission(PermissionType.PUBLISHER).withClientGroupName("exampleClientGroupName1").create();
    }
}

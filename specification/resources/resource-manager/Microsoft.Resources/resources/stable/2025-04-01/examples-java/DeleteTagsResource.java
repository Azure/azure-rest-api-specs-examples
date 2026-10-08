
/**
 * Samples for TagOperations DeleteAtScope.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01/DeleteTagsResource.json
     */
    /**
     * Sample code: Update tags on a resource.
     * 
     * @param manager Entry point to ResourceManager.
     */
    public static void updateTagsOnAResource(com.azure.resourcemanager.resources.ResourceManager manager) {
        manager.serviceClient().getTagOperations().deleteAtScope(
            "subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/my-resource-group/providers/myPRNameSpace/VM/myVm",
            com.azure.core.util.Context.NONE);
    }
}

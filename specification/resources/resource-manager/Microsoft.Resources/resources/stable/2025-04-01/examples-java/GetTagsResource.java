
/**
 * Samples for TagOperations GetAtScope.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01/GetTagsResource.json
     */
    /**
     * Sample code: Get tags on a resource.
     * 
     * @param manager Entry point to ResourceManager.
     */
    public static void getTagsOnAResource(com.azure.resourcemanager.resources.ResourceManager manager) {
        manager.serviceClient().getTagOperations().getAtScopeWithResponse(
            "subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/my-resource-group/providers/myPRNameSpace/VM/myVm",
            com.azure.core.util.Context.NONE);
    }
}

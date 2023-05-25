/** Samples for TagOperations DeleteAtScope. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2022-09-01/examples/DeleteTagsResource.json
     */
    /**
     * Sample code: Update tags on a resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateTagsOnAResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .serviceClient()
            .getTagOperations()
            .deleteAtScope(
                "subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/my-resource-group/providers/myPRNameSpace/VM/myVm",
                com.azure.core.util.Context.NONE);
    }
}

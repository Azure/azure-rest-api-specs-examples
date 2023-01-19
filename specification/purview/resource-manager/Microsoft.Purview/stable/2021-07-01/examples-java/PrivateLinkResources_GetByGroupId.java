/** Samples for PrivateLinkResources GetByGroupId. */
public final class Main {
    /*
     * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/PrivateLinkResources_GetByGroupId.json
     */
    /**
     * Sample code: PrivateLinkResources_GetByGroupId.
     *
     * @param manager Entry point to PurviewManager.
     */
    public static void privateLinkResourcesGetByGroupId(com.azure.resourcemanager.purview.PurviewManager manager) {
        manager
            .privateLinkResources()
            .getByGroupIdWithResponse("SampleResourceGroup", "account1", "group1", com.azure.core.util.Context.NONE);
    }
}

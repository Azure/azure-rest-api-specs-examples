/** Samples for PrivateLinkResources ListByAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/PrivateLinkResources_ListByAccount.json
     */
    /**
     * Sample code: PrivateLinkResources_ListByAccount.
     *
     * @param manager Entry point to PurviewManager.
     */
    public static void privateLinkResourcesListByAccount(com.azure.resourcemanager.purview.PurviewManager manager) {
        manager
            .privateLinkResources()
            .listByAccount("SampleResourceGroup", "account1", com.azure.core.util.Context.NONE);
    }
}

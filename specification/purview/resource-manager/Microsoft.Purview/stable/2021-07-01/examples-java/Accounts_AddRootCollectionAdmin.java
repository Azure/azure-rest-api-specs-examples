import com.azure.resourcemanager.purview.models.CollectionAdminUpdate;

/** Samples for Accounts AddRootCollectionAdmin. */
public final class Main {
    /*
     * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_AddRootCollectionAdmin.json
     */
    /**
     * Sample code: Accounts_AddRootCollectionAdmin.
     *
     * @param manager Entry point to PurviewManager.
     */
    public static void accountsAddRootCollectionAdmin(com.azure.resourcemanager.purview.PurviewManager manager) {
        manager
            .accounts()
            .addRootCollectionAdminWithResponse(
                "SampleResourceGroup",
                "account1",
                new CollectionAdminUpdate().withObjectId("7e8de0e7-2bfc-4e1f-9659-2a5785e4356f"),
                com.azure.core.util.Context.NONE);
    }
}

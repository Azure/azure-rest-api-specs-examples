/** Samples for Accounts CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_CreateOrUpdate.json
     */
    /**
     * Sample code: Accounts_CreateOrUpdate.
     *
     * @param manager Entry point to PurviewManager.
     */
    public static void accountsCreateOrUpdate(com.azure.resourcemanager.purview.PurviewManager manager) {
        manager
            .accounts()
            .define("account1")
            .withRegion("West US 2")
            .withExistingResourceGroup("SampleResourceGroup")
            .withManagedResourceGroupName("custom-rgname")
            .create();
    }
}

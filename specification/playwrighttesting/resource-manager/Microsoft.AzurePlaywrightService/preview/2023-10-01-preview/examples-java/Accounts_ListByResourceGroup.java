/** Samples for Accounts ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/playwrighttesting/resource-manager/Microsoft.AzurePlaywrightService/preview/2023-10-01-preview/examples/Accounts_ListByResourceGroup.json
     */
    /**
     * Sample code: Accounts_ListByResourceGroup.
     *
     * @param manager Entry point to PlaywrightTestingManager.
     */
    public static void accountsListByResourceGroup(
        com.azure.resourcemanager.playwrighttesting.PlaywrightTestingManager manager) {
        manager.accounts().listByResourceGroup("dummyrg", com.azure.core.util.Context.NONE);
    }
}

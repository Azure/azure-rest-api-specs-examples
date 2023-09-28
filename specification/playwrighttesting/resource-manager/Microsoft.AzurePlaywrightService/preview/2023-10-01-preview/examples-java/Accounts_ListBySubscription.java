/** Samples for Accounts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/playwrighttesting/resource-manager/Microsoft.AzurePlaywrightService/preview/2023-10-01-preview/examples/Accounts_ListBySubscription.json
     */
    /**
     * Sample code: Accounts_ListBySubscription.
     *
     * @param manager Entry point to PlaywrightTestingManager.
     */
    public static void accountsListBySubscription(
        com.azure.resourcemanager.playwrighttesting.PlaywrightTestingManager manager) {
        manager.accounts().list(com.azure.core.util.Context.NONE);
    }
}

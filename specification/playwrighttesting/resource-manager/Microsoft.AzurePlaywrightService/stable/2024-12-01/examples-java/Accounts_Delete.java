
/**
 * Samples for Accounts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/Accounts_Delete.json
     */
    /**
     * Sample code: Accounts_Delete.
     * 
     * @param manager Entry point to PlaywrightTestingManager.
     */
    public static void accountsDelete(com.azure.resourcemanager.playwrighttesting.PlaywrightTestingManager manager) {
        manager.accounts().delete("dummyrg", "myPlaywrightAccount", com.azure.core.util.Context.NONE);
    }
}

/** Samples for Accounts Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2022-04-01-preview/examples/Accounts/Accounts_Create.json
     */
    /**
     * Sample code: Creates or updates Account.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void createsOrUpdatesAccount(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.accounts().define("contoso").withRegion("westus2").withExistingResourceGroup("test-rg").create();
    }
}

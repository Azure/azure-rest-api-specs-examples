/** Samples for Accounts CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-05-01/examples/Accounts_CreateOrUpdate.json
     */
    /**
     * Sample code: Accounts_CreateOrUpdate.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void accountsCreateOrUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.accounts().define("account1").withRegion("eastus").withExistingResourceGroup("myRG").create();
    }
}

/** Samples for Accounts GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/playwrighttesting/resource-manager/Microsoft.AzurePlaywrightService/preview/2023-10-01-preview/examples/Accounts_Get.json
     */
    /**
     * Sample code: Accounts_Get.
     *
     * @param manager Entry point to PlaywrightTestingManager.
     */
    public static void accountsGet(com.azure.resourcemanager.playwrighttesting.PlaywrightTestingManager manager) {
        manager
            .accounts()
            .getByResourceGroupWithResponse("dummyrg", "myPlaywrightAccount", com.azure.core.util.Context.NONE);
    }
}

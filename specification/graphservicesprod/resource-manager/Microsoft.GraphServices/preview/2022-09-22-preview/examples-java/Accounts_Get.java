/** Samples for AccountOperation GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/graphservicesprod/resource-manager/Microsoft.GraphServices/preview/2022-09-22-preview/examples/Accounts_Get.json
     */
    /**
     * Sample code: Get accounts.
     *
     * @param manager Entry point to GraphServicesManager.
     */
    public static void getAccounts(com.azure.resourcemanager.graphservices.GraphServicesManager manager) {
        manager
            .accountOperations()
            .getByResourceGroupWithResponse(
                "testResourceGroupGRAM", "11111111-aaaa-1111-bbbb-111111111111", com.azure.core.util.Context.NONE);
    }
}

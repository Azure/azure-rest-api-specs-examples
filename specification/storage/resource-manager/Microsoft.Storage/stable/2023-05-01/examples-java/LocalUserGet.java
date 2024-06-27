
/**
 * Samples for LocalUsersOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/LocalUserGet.json
     */
    /**
     * Sample code: GetLocalUser.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getLocalUser(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getLocalUsersOperations().getWithResponse("res6977",
            "sto2527", "user1", com.azure.core.util.Context.NONE);
    }
}

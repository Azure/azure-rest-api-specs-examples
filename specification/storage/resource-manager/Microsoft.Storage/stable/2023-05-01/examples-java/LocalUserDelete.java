
/**
 * Samples for LocalUsersOperation Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/LocalUserDelete.json
     */
    /**
     * Sample code: DeleteLocalUser.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteLocalUser(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getLocalUsersOperations().deleteWithResponse("res6977",
            "sto2527", "user1", com.azure.core.util.Context.NONE);
    }
}

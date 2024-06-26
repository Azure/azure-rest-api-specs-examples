
/**
 * Samples for LocalUsersOperation ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/LocalUserListKeys.json
     */
    /**
     * Sample code: ListLocalUserKeys.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listLocalUserKeys(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getLocalUsersOperations().listKeysWithResponse("res6977",
            "sto2527", "user1", com.azure.core.util.Context.NONE);
    }
}

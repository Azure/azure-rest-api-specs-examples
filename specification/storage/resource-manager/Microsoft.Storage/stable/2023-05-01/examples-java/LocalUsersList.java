
/**
 * Samples for LocalUsersOperation List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/LocalUsersList.json
     */
    /**
     * Sample code: ListLocalUsers.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listLocalUsers(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getLocalUsersOperations().list("res6977", "sto2527", null,
            null, null, com.azure.core.util.Context.NONE);
    }
}

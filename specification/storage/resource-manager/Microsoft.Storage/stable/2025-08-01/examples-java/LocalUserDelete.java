
/**
 * Samples for LocalUsersOperation Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/LocalUserDelete.json
     */
    /**
     * Sample code: DeleteLocalUser.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void deleteLocalUser(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getLocalUsersOperations().deleteWithResponse("res6977", "sto2527", "user1",
            com.azure.core.util.Context.NONE);
    }
}

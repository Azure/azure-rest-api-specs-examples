
/**
 * Samples for LocalUsersOperation List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/LocalUsersList.json
     */
    /**
     * Sample code: ListLocalUsers.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void listLocalUsers(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getLocalUsersOperations().list("res6977", "sto2527", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}

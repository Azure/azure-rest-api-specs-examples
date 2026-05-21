
/**
 * Samples for LocalUsersOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/LocalUserGet.json
     */
    /**
     * Sample code: GetLocalUser.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void getLocalUser(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getLocalUsersOperations().getWithResponse("res6977", "sto2527", "user1",
            com.azure.core.util.Context.NONE);
    }
}

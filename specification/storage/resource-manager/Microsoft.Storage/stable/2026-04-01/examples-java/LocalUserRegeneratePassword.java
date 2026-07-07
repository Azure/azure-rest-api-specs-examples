
/**
 * Samples for LocalUsersOperation RegeneratePassword.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/LocalUserRegeneratePassword.json
     */
    /**
     * Sample code: RegenerateLocalUserPassword.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void regenerateLocalUserPassword(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getLocalUsersOperations().regeneratePasswordWithResponse("res6977", "sto2527", "user1",
            com.azure.core.util.Context.NONE);
    }
}

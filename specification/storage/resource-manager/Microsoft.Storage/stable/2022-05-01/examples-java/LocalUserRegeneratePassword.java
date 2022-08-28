import com.azure.core.util.Context;

/** Samples for LocalUsersOperation RegeneratePassword. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/LocalUserRegeneratePassword.json
     */
    /**
     * Sample code: RegenerateLocalUserPassword.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void regenerateLocalUserPassword(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getLocalUsersOperations()
            .regeneratePasswordWithResponse("res6977", "sto2527", "user1", Context.NONE);
    }
}

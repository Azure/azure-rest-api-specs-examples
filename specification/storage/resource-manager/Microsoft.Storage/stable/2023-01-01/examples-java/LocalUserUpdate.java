
import com.azure.resourcemanager.storage.fluent.models.LocalUserInner;

/** Samples for LocalUsersOperation CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/LocalUserUpdate.json
     */
    /**
     * Sample code: UpdateLocalUser.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateLocalUser(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getLocalUsersOperations().createOrUpdateWithResponse(
            "res6977", "sto2527", "user1", new LocalUserInner().withHomeDirectory("homedirectory2")
                .withHasSharedKey(false).withHasSshKey(false).withHasSshPassword(false),
            com.azure.core.util.Context.NONE);
    }
}


import com.azure.resourcemanager.storage.fluent.models.LocalUserInner;
import java.util.Arrays;

/**
 * Samples for LocalUsersOperation CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/LocalUserUpdate.json
     */
    /**
     * Sample code: UpdateLocalUser.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateLocalUser(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getLocalUsersOperations().createOrUpdateWithResponse(
            "res6977", "sto2527", "user1",
            new LocalUserInner().withHomeDirectory("homedirectory2").withHasSharedKey(false).withHasSshKey(false)
                .withHasSshPassword(false).withGroupId(3000).withAllowAclAuthorization(false)
                .withExtendedGroups(Arrays.asList(1001, 1005, 2005)).withIsNFSv3Enabled(true),
            com.azure.core.util.Context.NONE);
    }
}

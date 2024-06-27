
import com.azure.resourcemanager.storage.fluent.models.LocalUserInner;
import java.util.Arrays;

/**
 * Samples for LocalUsersOperation CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/LocalUserCreateNFSv3Enabled.
     * json
     */
    /**
     * Sample code: CreateNFSv3EnabledLocalUser.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createNFSv3EnabledLocalUser(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getLocalUsersOperations().createOrUpdateWithResponse(
            "res6977", "sto2527", "user1",
            new LocalUserInner().withExtendedGroups(Arrays.asList(1001, 1005, 2005)).withIsNFSv3Enabled(true),
            com.azure.core.util.Context.NONE);
    }
}

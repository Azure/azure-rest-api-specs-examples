
import com.azure.resourcemanager.storage.fluent.models.LocalUserInner;
import java.util.Arrays;

/**
 * Samples for LocalUsersOperation CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/LocalUserCreateNFSv3Enabled.json
     */
    /**
     * Sample code: CreateNFSv3EnabledLocalUser.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void createNFSv3EnabledLocalUser(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getLocalUsersOperations().createOrUpdateWithResponse("res6977", "sto2527", "user1",
            new LocalUserInner().withExtendedGroups(Arrays.asList(1001, 1005, 2005)).withIsNFSv3Enabled(true),
            com.azure.core.util.Context.NONE);
    }
}

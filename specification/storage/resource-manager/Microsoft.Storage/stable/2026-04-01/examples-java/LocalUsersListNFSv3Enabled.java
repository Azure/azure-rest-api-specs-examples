
import com.azure.resourcemanager.storage.models.ListLocalUserIncludeParam;

/**
 * Samples for LocalUsersOperation List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/LocalUsersListNFSv3Enabled.json
     */
    /**
     * Sample code: ListNFSv3EnabledLocalUsers.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void listNFSv3EnabledLocalUsers(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getLocalUsersOperations().list("res6977", "sto2527", null, null,
            ListLocalUserIncludeParam.NFSV3, com.azure.core.util.Context.NONE);
    }
}

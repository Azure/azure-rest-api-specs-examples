
import com.azure.resourcemanager.storage.models.ListLocalUserIncludeParam;

/**
 * Samples for LocalUsersOperation List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/LocalUsersListNFSv3Enabled.
     * json
     */
    /**
     * Sample code: ListNFSv3EnabledLocalUsers.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listNFSv3EnabledLocalUsers(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getLocalUsersOperations().list("res6977", "sto2527", null,
            null, ListLocalUserIncludeParam.NFSV3, com.azure.core.util.Context.NONE);
    }
}

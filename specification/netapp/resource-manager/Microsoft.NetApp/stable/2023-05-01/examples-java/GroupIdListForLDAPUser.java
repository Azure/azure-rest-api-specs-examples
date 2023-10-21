import com.azure.resourcemanager.netapp.models.GetGroupIdListForLdapUserRequest;

/** Samples for Volumes ListGetGroupIdListForLdapUser. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-05-01/examples/GroupIdListForLDAPUser.json
     */
    /**
     * Sample code: GetGroupIdListForUser.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void getGroupIdListForUser(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager
            .volumes()
            .listGetGroupIdListForLdapUser(
                "myRG",
                "account1",
                "pool1",
                "volume1",
                new GetGroupIdListForLdapUserRequest().withUsername("user1"),
                com.azure.core.util.Context.NONE);
    }
}

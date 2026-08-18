
/**
 * Samples for ExpressRouteLags MembersList.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ExpressRouteLagMemberList.json
     */
    /**
     * Sample code: List express route lag members.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listExpressRouteLagMembers(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteLags().membersList("rg1", "lagName", "linkName",
            com.azure.core.util.Context.NONE);
    }
}

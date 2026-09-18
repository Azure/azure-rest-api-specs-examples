
/**
 * Samples for ExpressRouteLags MembersGet.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/ExpressRouteLagMemberGet.json
     */
    /**
     * Sample code: Get express route lag member.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getExpressRouteLagMember(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteLags().membersGetWithResponse("rg1", "lagName", "linkName", "memberName",
            com.azure.core.util.Context.NONE);
    }
}

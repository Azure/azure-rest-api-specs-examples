
/**
 * Samples for ApplicationSecurityGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationSecurityGroupListAll.json
     */
    /**
     * Sample code: List all application security groups.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllApplicationSecurityGroups(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationSecurityGroups().list(com.azure.core.util.Context.NONE);
    }
}

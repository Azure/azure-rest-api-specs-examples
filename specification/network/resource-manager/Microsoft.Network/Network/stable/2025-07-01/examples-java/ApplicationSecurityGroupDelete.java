
/**
 * Samples for ApplicationSecurityGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationSecurityGroupDelete.json
     */
    /**
     * Sample code: Delete application security group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteApplicationSecurityGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationSecurityGroups().delete("rg1", "test-asg",
            com.azure.core.util.Context.NONE);
    }
}

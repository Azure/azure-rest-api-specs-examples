
/**
 * Samples for ApplicationSecurityGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationSecurityGroupGet.json
     */
    /**
     * Sample code: Get application security group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getApplicationSecurityGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationSecurityGroups().getByResourceGroupWithResponse("rg1", "test-asg",
            com.azure.core.util.Context.NONE);
    }
}

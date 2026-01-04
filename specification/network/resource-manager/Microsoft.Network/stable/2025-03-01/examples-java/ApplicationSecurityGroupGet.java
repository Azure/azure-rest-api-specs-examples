
/**
 * Samples for ApplicationSecurityGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ApplicationSecurityGroupGet.
     * json
     */
    /**
     * Sample code: Get application security group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getApplicationSecurityGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getApplicationSecurityGroups().getByResourceGroupWithResponse("rg1",
            "test-asg", com.azure.core.util.Context.NONE);
    }
}

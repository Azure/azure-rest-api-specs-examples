
/**
 * Samples for ApplicationSecurityGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * ApplicationSecurityGroupDelete.json
     */
    /**
     * Sample code: Delete application security group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteApplicationSecurityGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getApplicationSecurityGroups().delete("rg1", "test-asg",
            com.azure.core.util.Context.NONE);
    }
}

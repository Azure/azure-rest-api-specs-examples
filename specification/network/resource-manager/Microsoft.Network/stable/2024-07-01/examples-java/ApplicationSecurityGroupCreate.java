
import com.azure.resourcemanager.network.fluent.models.ApplicationSecurityGroupInner;

/**
 * Samples for ApplicationSecurityGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * ApplicationSecurityGroupCreate.json
     */
    /**
     * Sample code: Create application security group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createApplicationSecurityGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getApplicationSecurityGroups().createOrUpdate("rg1", "test-asg",
            new ApplicationSecurityGroupInner().withLocation("westus"), com.azure.core.util.Context.NONE);
    }
}

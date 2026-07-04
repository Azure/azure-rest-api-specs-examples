
import com.azure.resourcemanager.network.fluent.models.ApplicationSecurityGroupInner;

/**
 * Samples for ApplicationSecurityGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationSecurityGroupCreate.json
     */
    /**
     * Sample code: Create application security group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createApplicationSecurityGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationSecurityGroups().createOrUpdate("rg1", "test-asg",
            new ApplicationSecurityGroupInner().withLocation("westus"), com.azure.core.util.Context.NONE);
    }
}

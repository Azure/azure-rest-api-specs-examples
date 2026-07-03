
import com.azure.resourcemanager.network.fluent.models.NetworkSecurityGroupInner;

/**
 * Samples for NetworkSecurityGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkSecurityGroupCreate.json
     */
    /**
     * Sample code: Create network security group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createNetworkSecurityGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityGroups().createOrUpdate("rg1", "testnsg",
            new NetworkSecurityGroupInner().withLocation("eastus"), com.azure.core.util.Context.NONE);
    }
}

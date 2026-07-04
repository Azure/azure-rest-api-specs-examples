
import com.azure.resourcemanager.network.fluent.models.SubnetInner;

/**
 * Samples for Subnets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/SubnetCreateWithSharingScope.json
     */
    /**
     * Sample code: Create subnet with sharing scope.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createSubnetWithSharingScope(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSubnets().createOrUpdate("subnet-test", "vnetname", "subnet1",
            new SubnetInner().withAddressPrefix("10.0.0.0/16"), com.azure.core.util.Context.NONE);
    }
}

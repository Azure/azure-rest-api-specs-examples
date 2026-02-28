
import com.azure.resourcemanager.network.fluent.models.SubnetInner;

/**
 * Samples for Subnets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/SubnetCreate.json
     */
    /**
     * Sample code: Create subnet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createSubnet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSubnets().createOrUpdate("subnet-test", "vnetname", "subnet1",
            new SubnetInner().withAddressPrefix("10.0.0.0/16"), com.azure.core.util.Context.NONE);
    }
}

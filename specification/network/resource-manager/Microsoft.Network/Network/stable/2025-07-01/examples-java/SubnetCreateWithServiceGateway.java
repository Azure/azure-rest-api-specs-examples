
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.SubnetInner;

/**
 * Samples for Subnets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/SubnetCreateWithServiceGateway.json
     */
    /**
     * Sample code: Create Subnet with service gateway.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createSubnetWithServiceGateway(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSubnets().createOrUpdate("subnet-test", "vnetname", "subnet1",
            new SubnetInner().withAddressPrefix("10.0.0.0/16").withServiceGateway(new SubResource().withId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/serviceGateways/SG1")),
            com.azure.core.util.Context.NONE);
    }
}

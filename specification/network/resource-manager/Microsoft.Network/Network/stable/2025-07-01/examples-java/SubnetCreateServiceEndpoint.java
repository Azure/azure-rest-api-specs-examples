
import com.azure.resourcemanager.network.fluent.models.SubnetInner;
import com.azure.resourcemanager.network.models.ServiceEndpointPropertiesFormat;
import java.util.Arrays;

/**
 * Samples for Subnets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/SubnetCreateServiceEndpoint.json
     */
    /**
     * Sample code: Create subnet with service endpoints.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createSubnetWithServiceEndpoints(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSubnets().createOrUpdate("subnet-test", "vnetname", "subnet1",
            new SubnetInner().withAddressPrefix("10.0.0.0/16").withServiceEndpoints(
                Arrays.asList(new ServiceEndpointPropertiesFormat().withService("Microsoft.Storage"))),
            com.azure.core.util.Context.NONE);
    }
}

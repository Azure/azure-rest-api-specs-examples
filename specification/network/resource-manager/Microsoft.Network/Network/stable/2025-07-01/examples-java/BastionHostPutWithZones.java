
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.BastionHostInner;
import com.azure.resourcemanager.network.models.BastionHostIpConfiguration;
import java.util.Arrays;

/**
 * Samples for BastionHosts CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/BastionHostPutWithZones.json
     */
    /**
     * Sample code: Create Bastion Host With Zones.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createBastionHostWithZones(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getBastionHosts().createOrUpdate("rg1", "bastionhosttenant",
            new BastionHostInner().withIpConfigurations(Arrays.asList(new BastionHostIpConfiguration()
                .withName("bastionHostIpConfiguration")
                .withSubnet(new SubResource().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/BastionHostSubnet"))
                .withPublicIpAddress(new SubResource().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName")))),
            com.azure.core.util.Context.NONE);
    }
}

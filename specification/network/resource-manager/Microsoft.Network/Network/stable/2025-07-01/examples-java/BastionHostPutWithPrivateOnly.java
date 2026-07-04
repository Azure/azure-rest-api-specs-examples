
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.BastionHostInner;
import com.azure.resourcemanager.network.models.BastionHostIpConfiguration;
import java.util.Arrays;

/**
 * Samples for BastionHosts CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/BastionHostPutWithPrivateOnly.json
     */
    /**
     * Sample code: Create Bastion Host With Private Only.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createBastionHostWithPrivateOnly(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getBastionHosts().createOrUpdate("rg1", "bastionhosttenant", new BastionHostInner()
            .withIpConfigurations(Arrays.asList(new BastionHostIpConfiguration().withName("bastionHostIpConfiguration")
                .withSubnet(new SubResource().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/BastionHostSubnet"))))
            .withEnablePrivateOnlyBastion(true), com.azure.core.util.Context.NONE);
    }
}

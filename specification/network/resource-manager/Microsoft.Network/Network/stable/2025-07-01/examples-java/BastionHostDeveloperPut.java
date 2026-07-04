
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.BastionHostInner;
import com.azure.resourcemanager.network.models.BastionHostPropertiesFormatNetworkAcls;
import com.azure.resourcemanager.network.models.IpRule;
import java.util.Arrays;

/**
 * Samples for BastionHosts CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/BastionHostDeveloperPut.json
     */
    /**
     * Sample code: Create Developer Bastion Host.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createDeveloperBastionHost(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getBastionHosts().createOrUpdate("rg2", "bastionhostdeveloper",
            new BastionHostInner().withIpConfigurations(Arrays.asList()).withVirtualNetwork(new SubResource().withId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg2/providers/Microsoft.Network/virtualNetworks/vnet2"))
                .withNetworkAcls(new BastionHostPropertiesFormatNetworkAcls()
                    .withIpRules(Arrays.asList(new IpRule().withAddressPrefix("1.1.1.1/16")))),
            com.azure.core.util.Context.NONE);
    }
}


import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.NatGatewayInner;
import com.azure.resourcemanager.network.models.Nat64State;
import com.azure.resourcemanager.network.models.NatGatewaySku;
import com.azure.resourcemanager.network.models.NatGatewaySkuName;
import java.util.Arrays;

/**
 * Samples for NatGateways CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NatGatewayWithNat64CreateOrUpdate.json
     */
    /**
     * Sample code: Create nat gateway with nat64.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createNatGatewayWithNat64(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNatGateways().createOrUpdate("rg1", "test-natgateway", new NatGatewayInner()
            .withLocation("westus").withSku(new NatGatewaySku().withName(NatGatewaySkuName.STANDARD))
            .withPublicIpAddresses(Arrays.asList(new SubResource().withId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/PublicIpAddress1")))
            .withPublicIpPrefixes(Arrays.asList(new SubResource().withId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/PublicIpPrefix1")))
            .withNat64(Nat64State.ENABLED), com.azure.core.util.Context.NONE);
    }
}

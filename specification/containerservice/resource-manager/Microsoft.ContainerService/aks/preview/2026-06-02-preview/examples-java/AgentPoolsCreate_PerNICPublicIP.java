
import com.azure.resourcemanager.containerservice.fluent.models.AgentPoolInner;
import com.azure.resourcemanager.containerservice.models.AgentPoolNICPublicIPAddressConfiguration;
import com.azure.resourcemanager.containerservice.models.AgentPoolNICPublicIPAddressVersion;
import com.azure.resourcemanager.containerservice.models.AgentPoolNetworkInterface;
import com.azure.resourcemanager.containerservice.models.AgentPoolNetworkInterfaceType;
import com.azure.resourcemanager.containerservice.models.AgentPoolNetworkProfile;
import com.azure.resourcemanager.containerservice.models.IpTag;
import com.azure.resourcemanager.containerservice.models.OSType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AgentPools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-02-preview/AgentPoolsCreate_PerNICPublicIP.json
     */
    /**
     * Sample code: Create Agent Pool with per-NIC public IP configuration.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void createAgentPoolWithPerNICPublicIPConfiguration(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentPools().createOrUpdate("rg1", "clustername1", "agentpool1", new AgentPoolInner()
            .withCount(3).withVmSize("Standard_D8s_v3").withOsType(OSType.LINUX).withOrchestratorVersion("")
            .withNetworkProfile(new AgentPoolNetworkProfile().withSecondaryNetworkInterfaces(Arrays.asList(
                new AgentPoolNetworkInterface().withType(AgentPoolNetworkInterfaceType.STANDARD).withVnetSubnetId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/myVNet/subnets/secondary-subnet-1")
                    .withEnableAcceleratedNetworking(true)
                    .withPublicIPAddressConfiguration(
                        new AgentPoolNICPublicIPAddressConfiguration().withPublicIPAddressVersion(
                            AgentPoolNICPublicIPAddressVersion.IPV4)
                            .withIpTags(Arrays.asList(new IpTag().withIpTagType("FirstPartyUsage").withTag("teams")))),
                new AgentPoolNetworkInterface().withType(AgentPoolNetworkInterfaceType.STANDARD).withVnetSubnetId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/myVNet/subnets/secondary-subnet-2")
                    .withEnableAcceleratedNetworking(true)
                    .withPublicIPAddressConfiguration(new AgentPoolNICPublicIPAddressConfiguration()
                        .withPublicIPAddressVersion(AgentPoolNICPublicIPAddressVersion.IPV4).withPublicIPPrefixID(
                            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/myPrefix"))))),
            null, null, com.azure.core.util.Context.NONE);
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}

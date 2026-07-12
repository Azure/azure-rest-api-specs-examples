
import com.azure.resourcemanager.batch.models.DeploymentConfiguration;
import com.azure.resourcemanager.batch.models.FixedScaleSettings;
import com.azure.resourcemanager.batch.models.IPFamily;
import com.azure.resourcemanager.batch.models.ImageReference;
import com.azure.resourcemanager.batch.models.InboundEndpointProtocol;
import com.azure.resourcemanager.batch.models.InboundNatPool;
import com.azure.resourcemanager.batch.models.NetworkConfiguration;
import com.azure.resourcemanager.batch.models.NetworkSecurityGroupRule;
import com.azure.resourcemanager.batch.models.NetworkSecurityGroupRuleAccess;
import com.azure.resourcemanager.batch.models.PoolEndpointConfiguration;
import com.azure.resourcemanager.batch.models.PublicIpAddressConfiguration;
import com.azure.resourcemanager.batch.models.ScaleSettings;
import com.azure.resourcemanager.batch.models.VirtualMachineConfiguration;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Pool Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/PoolCreate_DualStackNetworking.json
     */
    /**
     * Sample code: CreatePool - dual stack networking.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void createPoolDualStackNetworking(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.pools().define("dualstackpool").withExistingBatchAccount("default-azurebatch", "exampleacc")
            .withVmSize(
                "Standard_D4ds_v5")
            .withDeploymentConfiguration(
                new DeploymentConfiguration()
                    .withVirtualMachineConfiguration(new VirtualMachineConfiguration()
                        .withImageReference(new ImageReference().withPublisher("Canonical")
                            .withOffer("ubuntu-24_04-lts").withSku("server"))
                        .withNodeAgentSkuId("batch.node.ubuntu 24.04")))
            .withScaleSettings(new ScaleSettings()
                .withFixedScale(new FixedScaleSettings().withTargetDedicatedNodes(1).withTargetLowPriorityNodes(0)))
            .withNetworkConfiguration(new NetworkConfiguration()
                .withEndpointConfiguration(new PoolEndpointConfiguration().withInboundNatPools(
                    Arrays.asList(new InboundNatPool().withName("sshpool").withProtocol(InboundEndpointProtocol.TCP)
                        .withBackendPort(22).withFrontendPortRangeStart(40000).withFrontendPortRangeEnd(40500)
                        .withNetworkSecurityGroupRules(Arrays.asList(new NetworkSecurityGroupRule().withPriority(1000)
                            .withAccess(NetworkSecurityGroupRuleAccess.ALLOW).withSourceAddressPrefix("*")
                            .withSourcePortRanges(Arrays.asList("*")))))))
                .withPublicIpAddressConfiguration(
                    new PublicIpAddressConfiguration().withIpFamilies(Arrays.asList(IPFamily.IPV4, IPFamily.IPV6))))
            .create();
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

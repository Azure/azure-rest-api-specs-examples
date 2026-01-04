
import com.azure.resourcemanager.containerservice.fluent.models.ManagedClusterInner;
import com.azure.resourcemanager.containerservice.models.AgentPoolMode;
import com.azure.resourcemanager.containerservice.models.AgentPoolType;
import com.azure.resourcemanager.containerservice.models.ContainerServiceLinuxProfile;
import com.azure.resourcemanager.containerservice.models.ContainerServiceNetworkProfile;
import com.azure.resourcemanager.containerservice.models.ContainerServiceSshConfiguration;
import com.azure.resourcemanager.containerservice.models.ContainerServiceSshPublicKey;
import com.azure.resourcemanager.containerservice.models.LoadBalancerSku;
import com.azure.resourcemanager.containerservice.models.ManagedClusterAgentPoolProfile;
import com.azure.resourcemanager.containerservice.models.ManagedClusterApiServerAccessProfile;
import com.azure.resourcemanager.containerservice.models.ManagedClusterLoadBalancerProfile;
import com.azure.resourcemanager.containerservice.models.ManagedClusterLoadBalancerProfileManagedOutboundIPs;
import com.azure.resourcemanager.containerservice.models.ManagedClusterPropertiesAutoScalerProfile;
import com.azure.resourcemanager.containerservice.models.ManagedClusterServicePrincipalProfile;
import com.azure.resourcemanager.containerservice.models.ManagedClusterSku;
import com.azure.resourcemanager.containerservice.models.ManagedClusterSkuName;
import com.azure.resourcemanager.containerservice.models.ManagedClusterSkuTier;
import com.azure.resourcemanager.containerservice.models.ManagedClusterWindowsProfile;
import com.azure.resourcemanager.containerservice.models.OSType;
import com.azure.resourcemanager.containerservice.models.OutboundType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ManagedClusters CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * ManagedClustersCreate_DisableRunCommand.json
     */
    /**
     * Sample code: Create Managed Cluster with RunCommand disabled.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createManagedClusterWithRunCommandDisabled(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getManagedClusters().createOrUpdate("rg1", "clustername1",
            new ManagedClusterInner().withLocation("location1").withTags(mapOf("archv2", "", "tier", "production"))
                .withSku(new ManagedClusterSku()
                    .withName(ManagedClusterSkuName.fromString("Basic")).withTier(ManagedClusterSkuTier.FREE))
                .withKubernetesVersion("").withDnsPrefix("dnsprefix1")
                .withAgentPoolProfiles(Arrays.asList(new ManagedClusterAgentPoolProfile().withCount(3)
                    .withVmSize("Standard_DS2_v2").withOsType(OSType.LINUX)
                    .withType(AgentPoolType.VIRTUAL_MACHINE_SCALE_SETS).withMode(AgentPoolMode.SYSTEM)
                    .withEnableNodePublicIp(true).withEnableEncryptionAtHost(true).withName("nodepool1")))
                .withLinuxProfile(new ContainerServiceLinuxProfile().withAdminUsername("azureuser")
                    .withSsh(new ContainerServiceSshConfiguration().withPublicKeys(
                        Arrays.asList(new ContainerServiceSshPublicKey().withKeyData("fakeTokenPlaceholder")))))
                .withWindowsProfile(new ManagedClusterWindowsProfile().withAdminUsername("azureuser")
                    .withAdminPassword("fakeTokenPlaceholder"))
                .withServicePrincipalProfile(new ManagedClusterServicePrincipalProfile()
                    .withClientId("clientid").withSecret("fakeTokenPlaceholder"))
                .withAddonProfiles(mapOf()).withEnableRbac(true)
                .withNetworkProfile(new ContainerServiceNetworkProfile().withOutboundType(OutboundType.LOAD_BALANCER)
                    .withLoadBalancerSku(LoadBalancerSku.STANDARD)
                    .withLoadBalancerProfile(new ManagedClusterLoadBalancerProfile().withManagedOutboundIPs(
                        new ManagedClusterLoadBalancerProfileManagedOutboundIPs().withCount(2))))
                .withAutoScalerProfile(new ManagedClusterPropertiesAutoScalerProfile().withScanInterval("20s")
                    .withScaleDownDelayAfterAdd("15m"))
                .withApiServerAccessProfile(new ManagedClusterApiServerAccessProfile().withDisableRunCommand(true)),
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

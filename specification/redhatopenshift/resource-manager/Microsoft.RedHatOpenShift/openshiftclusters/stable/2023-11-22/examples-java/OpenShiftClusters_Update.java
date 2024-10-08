
import com.azure.resourcemanager.redhatopenshift.models.ApiServerProfile;
import com.azure.resourcemanager.redhatopenshift.models.ClusterProfile;
import com.azure.resourcemanager.redhatopenshift.models.ConsoleProfile;
import com.azure.resourcemanager.redhatopenshift.models.EncryptionAtHost;
import com.azure.resourcemanager.redhatopenshift.models.FipsValidatedModules;
import com.azure.resourcemanager.redhatopenshift.models.IngressProfile;
import com.azure.resourcemanager.redhatopenshift.models.LoadBalancerProfile;
import com.azure.resourcemanager.redhatopenshift.models.ManagedOutboundIPs;
import com.azure.resourcemanager.redhatopenshift.models.MasterProfile;
import com.azure.resourcemanager.redhatopenshift.models.NetworkProfile;
import com.azure.resourcemanager.redhatopenshift.models.OpenShiftCluster;
import com.azure.resourcemanager.redhatopenshift.models.PreconfiguredNsg;
import com.azure.resourcemanager.redhatopenshift.models.ServicePrincipalProfile;
import com.azure.resourcemanager.redhatopenshift.models.Visibility;
import com.azure.resourcemanager.redhatopenshift.models.WorkerProfile;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for OpenShiftClusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/
     * examples/OpenShiftClusters_Update.json
     */
    /**
     * Sample code: Updates a OpenShift cluster with the specified subscription, resource group and resource name.
     * 
     * @param manager Entry point to RedHatOpenShiftManager.
     */
    public static void updatesAOpenShiftClusterWithTheSpecifiedSubscriptionResourceGroupAndResourceName(
        com.azure.resourcemanager.redhatopenshift.RedHatOpenShiftManager manager) {
        OpenShiftCluster resource = manager.openShiftClusters()
            .getByResourceGroupWithResponse("resourceGroup", "resourceName", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("key", "fakeTokenPlaceholder"))
            .withClusterProfile(
                new ClusterProfile().withPullSecret("fakeTokenPlaceholder").withDomain("cluster.location.aroapp.io")
                    .withResourceGroupId("/subscriptions/subscriptionId/resourceGroups/clusterResourceGroup")
                    .withFipsValidatedModules(FipsValidatedModules.ENABLED))
            .withConsoleProfile(new ConsoleProfile())
            .withServicePrincipalProfile(
                new ServicePrincipalProfile().withClientId("clientId").withClientSecret("fakeTokenPlaceholder"))
            .withNetworkProfile(new NetworkProfile().withPodCidr("10.128.0.0/14").withServiceCidr("172.30.0.0/16")
                .withLoadBalancerProfile(
                    new LoadBalancerProfile().withManagedOutboundIps(new ManagedOutboundIPs().withCount(1)))
                .withPreconfiguredNsg(PreconfiguredNsg.DISABLED))
            .withMasterProfile(new MasterProfile().withVmSize("Standard_D8s_v3").withSubnetId(
                "/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/master")
                .withEncryptionAtHost(EncryptionAtHost.ENABLED))
            .withWorkerProfiles(Arrays.asList(new WorkerProfile().withName("worker").withVmSize("Standard_D2s_v3")
                .withDiskSizeGB(128)
                .withSubnetId(
                    "/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/worker")
                .withCount(3)))
            .withApiserverProfile(new ApiServerProfile().withVisibility(Visibility.PUBLIC)).withIngressProfiles(
                Arrays.asList(new IngressProfile().withName("default").withVisibility(Visibility.PUBLIC)))
            .apply();
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

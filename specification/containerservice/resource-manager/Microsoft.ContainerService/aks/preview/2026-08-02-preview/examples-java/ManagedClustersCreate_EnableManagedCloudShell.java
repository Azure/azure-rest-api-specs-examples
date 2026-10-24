
import com.azure.resourcemanager.containerservice.fluent.models.ManagedClusterInner;
import com.azure.resourcemanager.containerservice.models.CloudShellProfile;
import com.azure.resourcemanager.containerservice.models.CloudShellProfileMode;
import com.azure.resourcemanager.containerservice.models.CloudShellStorageProfile;
import com.azure.resourcemanager.containerservice.models.CloudShellStorageProfileMode;
import com.azure.resourcemanager.containerservice.models.ClusterApiConnectivityProfile;
import com.azure.resourcemanager.containerservice.models.ContainerServiceNetworkProfile;
import com.azure.resourcemanager.containerservice.models.ManagedClusterApiServerAccessProfile;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ManagedClusters CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-02-preview/ManagedClustersCreate_EnableManagedCloudShell.json
     */
    /**
     * Sample code: Create Managed Private Cluster With AKS-Managed Cloud Shell Resources.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void createManagedPrivateClusterWithAKSManagedCloudShellResources(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().createOrUpdate("rg1", "clustername1",
            new ManagedClusterInner().withLocation("location1").withNetworkProfile(new ContainerServiceNetworkProfile()
                .withClusterApiConnectivityProfile(new ClusterApiConnectivityProfile()
                    .withCloudShell(new CloudShellProfile().withMode(CloudShellProfileMode.MANAGED)
                        .withStorage(new CloudShellStorageProfile().withMode(CloudShellStorageProfileMode.MANAGED)))))
                .withApiServerAccessProfile(new ManagedClusterApiServerAccessProfile().withEnablePrivateCluster(true)),
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


import com.azure.resourcemanager.servicefabricmanagedclusters.models.ClusterUpgradeCadence;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.ClusterUpgradeMode;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.SettingsParameterDescription;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.SettingsSectionDescription;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.Sku;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.SkuName;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ManagedClusters CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/ManagedClusterPutOperation_example_min.json
     */
    /**
     * Sample code: Put a cluster with minimum parameters.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void putAClusterWithMinimumParameters(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.managedClusters().define("myCluster").withRegion("eastus").withExistingResourceGroup("resRg")
            .withSku(new Sku().withName(SkuName.BASIC)).withDnsName("myCluster").withAdminUsername("vmadmin")
            .withAdminPassword("{vm-password}")
            .withFabricSettings(
                Arrays.asList(new SettingsSectionDescription().withName("ManagedIdentityTokenService").withParameters(
                    Arrays.asList(new SettingsParameterDescription().withName("IsEnabled").withValue("true")))))
            .withClusterUpgradeMode(ClusterUpgradeMode.AUTOMATIC).withClusterUpgradeCadence(ClusterUpgradeCadence.WAVE1)
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

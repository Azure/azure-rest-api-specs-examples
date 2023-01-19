import com.azure.resourcemanager.servicefabric.models.DiagnosticsStorageAccountConfig;
import com.azure.resourcemanager.servicefabric.models.DurabilityLevel;
import com.azure.resourcemanager.servicefabric.models.EndpointRangeDescription;
import com.azure.resourcemanager.servicefabric.models.NodeTypeDescription;
import com.azure.resourcemanager.servicefabric.models.ReliabilityLevel;
import com.azure.resourcemanager.servicefabric.models.SettingsParameterDescription;
import com.azure.resourcemanager.servicefabric.models.SettingsSectionDescription;
import com.azure.resourcemanager.servicefabric.models.UpgradeMode;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Clusters CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterPutOperation_example_min.json
     */
    /**
     * Sample code: Put a cluster with minimum parameters.
     *
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void putAClusterWithMinimumParameters(
        com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager
            .clusters()
            .define("myCluster")
            .withRegion("eastus")
            .withExistingResourceGroup("resRg")
            .withTags(mapOf())
            .withDiagnosticsStorageAccountConfig(
                new DiagnosticsStorageAccountConfig()
                    .withStorageAccountName("diag")
                    .withProtectedAccountKeyName("fakeTokenPlaceholder")
                    .withBlobEndpoint("https://diag.blob.core.windows.net/")
                    .withQueueEndpoint("https://diag.queue.core.windows.net/")
                    .withTableEndpoint("https://diag.table.core.windows.net/"))
            .withFabricSettings(
                Arrays
                    .asList(
                        new SettingsSectionDescription()
                            .withName("UpgradeService")
                            .withParameters(
                                Arrays
                                    .asList(
                                        new SettingsParameterDescription()
                                            .withName("AppPollIntervalInSeconds")
                                            .withValue("60")))))
            .withManagementEndpoint("http://myCluster.eastus.cloudapp.azure.com:19080")
            .withNodeTypes(
                Arrays
                    .asList(
                        new NodeTypeDescription()
                            .withName("nt1vm")
                            .withClientConnectionEndpointPort(19000)
                            .withHttpGatewayEndpointPort(19007)
                            .withDurabilityLevel(DurabilityLevel.BRONZE)
                            .withApplicationPorts(
                                new EndpointRangeDescription().withStartPort(20000).withEndPort(30000))
                            .withEphemeralPorts(new EndpointRangeDescription().withStartPort(49000).withEndPort(64000))
                            .withIsPrimary(true)
                            .withVmInstanceCount(5)))
            .withReliabilityLevel(ReliabilityLevel.SILVER)
            .withUpgradeMode(UpgradeMode.AUTOMATIC)
            .create();
    }

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

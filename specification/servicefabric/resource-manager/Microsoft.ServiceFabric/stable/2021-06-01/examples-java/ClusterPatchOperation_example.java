
import com.azure.resourcemanager.servicefabric.models.Cluster;
import com.azure.resourcemanager.servicefabric.models.ClusterUpgradeCadence;
import com.azure.resourcemanager.servicefabric.models.DurabilityLevel;
import com.azure.resourcemanager.servicefabric.models.EndpointRangeDescription;
import com.azure.resourcemanager.servicefabric.models.NodeTypeDescription;
import com.azure.resourcemanager.servicefabric.models.ReliabilityLevel;
import com.azure.resourcemanager.servicefabric.models.UpgradeMode;
import java.time.OffsetDateTime;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Clusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/
     * ClusterPatchOperation_example.json
     */
    /**
     * Sample code: Patch a cluster.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void patchACluster(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        Cluster resource = manager.clusters()
            .getByResourceGroupWithResponse("resRg", "myCluster", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("a", "b")).withEventStoreServiceEnabled(true)
            .withNodeTypes(
                Arrays
                    .asList(
                        new NodeTypeDescription().withName("nt1vm").withClientConnectionEndpointPort(19000)
                            .withHttpGatewayEndpointPort(19007).withDurabilityLevel(DurabilityLevel.BRONZE)
                            .withApplicationPorts(
                                new EndpointRangeDescription().withStartPort(20000).withEndPort(30000))
                            .withEphemeralPorts(new EndpointRangeDescription().withStartPort(49000).withEndPort(64000))
                            .withIsPrimary(true).withVmInstanceCount(5),
                        new NodeTypeDescription().withName("testnt1").withClientConnectionEndpointPort(0)
                            .withHttpGatewayEndpointPort(0).withDurabilityLevel(DurabilityLevel.BRONZE)
                            .withApplicationPorts(new EndpointRangeDescription().withStartPort(1000).withEndPort(2000))
                            .withEphemeralPorts(new EndpointRangeDescription().withStartPort(3000).withEndPort(4000))
                            .withIsPrimary(false).withVmInstanceCount(3)))
            .withReliabilityLevel(ReliabilityLevel.BRONZE).withUpgradeMode(UpgradeMode.AUTOMATIC)
            .withUpgradeWave(ClusterUpgradeCadence.fromString("Wave"))
            .withUpgradePauseStartTimestampUtc(OffsetDateTime.parse("2021-06-21T22:00:00Z"))
            .withUpgradePauseEndTimestampUtc(OffsetDateTime.parse("2021-06-25T22:00:00Z")).apply();
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

import com.azure.core.util.Context;
import com.azure.resourcemanager.containerservice.fluent.models.SnapshotInner;
import com.azure.resourcemanager.containerservice.models.CreationData;
import java.util.HashMap;
import java.util.Map;

/** Samples for Snapshots CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-04-01/examples/SnapshotsCreate.json
     */
    /**
     * Sample code: Create/Update Snapshot.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createUpdateSnapshot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getSnapshots()
            .createOrUpdateWithResponse(
                "rg1",
                "snapshot1",
                new SnapshotInner()
                    .withLocation("westus")
                    .withTags(mapOf("key1", "val1", "key2", "val2"))
                    .withCreationData(
                        new CreationData()
                            .withSourceResourceId(
                                "/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/cluster1/agentPools/pool0")),
                Context.NONE);
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

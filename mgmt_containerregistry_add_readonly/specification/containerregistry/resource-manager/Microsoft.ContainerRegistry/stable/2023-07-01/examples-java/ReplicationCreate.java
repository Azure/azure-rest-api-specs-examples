import com.azure.resourcemanager.containerregistry.fluent.models.ReplicationInner;
import java.util.HashMap;
import java.util.Map;

/** Samples for Replications Create. */
public final class Main {
    /*
     * x-ms-original-file: mgmt_containerregistry_add_readonly/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2023-07-01/examples/ReplicationCreate.json
     */
    /**
     * Sample code: ReplicationCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void replicationCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getReplications()
            .create(
                "myResourceGroup",
                "myRegistry",
                "myReplication",
                new ReplicationInner().withLocation("eastus").withTags(mapOf("key", "fakeTokenPlaceholder")),
                com.azure.core.util.Context.NONE);
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

import com.azure.core.util.Context;
import com.azure.resourcemanager.containerregistry.models.ReplicationUpdateParameters;
import java.util.HashMap;
import java.util.Map;

/** Samples for Replications Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2021-09-01/examples/ReplicationUpdate.json
     */
    /**
     * Sample code: ReplicationUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void replicationUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getReplications()
            .update(
                "myResourceGroup",
                "myRegistry",
                "myReplication",
                new ReplicationUpdateParameters().withTags(mapOf("key", "value")),
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

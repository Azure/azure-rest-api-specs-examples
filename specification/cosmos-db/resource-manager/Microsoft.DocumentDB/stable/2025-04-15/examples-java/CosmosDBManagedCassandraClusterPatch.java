
import com.azure.resourcemanager.cosmos.fluent.models.ClusterResourceInner;
import com.azure.resourcemanager.cosmos.models.AuthenticationMethod;
import com.azure.resourcemanager.cosmos.models.Certificate;
import com.azure.resourcemanager.cosmos.models.ClusterResourceProperties;
import com.azure.resourcemanager.cosmos.models.SeedNode;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for CassandraClusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/
     * CosmosDBManagedCassandraClusterPatch.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraClusterPatch.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBManagedCassandraClusterPatch(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getCassandraClusters()
            .update("cassandra-prod-rg", "cassandra-prod", new ClusterResourceInner().withTags(mapOf("owner", "mike"))
                .withProperties(new ClusterResourceProperties().withAuthenticationMethod(AuthenticationMethod.NONE)
                    .withExternalGossipCertificates(Arrays.asList(new Certificate().withPem(
                        "-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----")))
                    .withExternalSeedNodes(Arrays.asList(new SeedNode().withIpAddress("10.52.221.2"),
                        new SeedNode().withIpAddress("10.52.221.3"), new SeedNode().withIpAddress("10.52.221.4")))
                    .withHoursBetweenBackups(12)),
                com.azure.core.util.Context.NONE);
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

import java.util.HashMap;
import java.util.Map;

/** Samples for Clusters Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ClusterCreateReadReplica.json
     */
    /**
     * Sample code: Create a new cluster as a read replica.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void createANewClusterAsAReadReplica(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager
            .clusters()
            .define("testcluster")
            .withRegion("westus")
            .withExistingResourceGroup("TestGroup")
            .withSourceResourceId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/sourcecluster")
            .withSourceLocation("westus")
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

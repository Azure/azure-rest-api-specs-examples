import java.util.HashMap;
import java.util.Map;

/** Samples for Clusters Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ClusterCreateBurstablev1.json
     */
    /**
     * Sample code: Create a new single node Burstable 1 vCore cluster.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void createANewSingleNodeBurstable1VCoreCluster(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager
            .clusters()
            .define("testcluster-burstablev1")
            .withRegion("westus")
            .withExistingResourceGroup("TestGroup")
            .withTags(mapOf("owner", "JohnDoe"))
            .withAdministratorLoginPassword("password")
            .withPostgresqlVersion("15")
            .withCitusVersion("11.3")
            .withPreferredPrimaryZone("1")
            .withEnableShardsOnCoordinator(true)
            .withEnableHa(false)
            .withCoordinatorServerEdition("BurstableMemoryOptimized")
            .withCoordinatorStorageQuotaInMb(131072)
            .withCoordinatorVCores(1)
            .withCoordinatorEnablePublicIpAccess(true)
            .withNodeCount(0)
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


import com.azure.resourcemanager.cosmosdbforpostgresql.models.DataEncryption;
import com.azure.resourcemanager.cosmosdbforpostgresql.models.DataEncryptionType;
import com.azure.resourcemanager.cosmosdbforpostgresql.models.IdentityProperties;
import com.azure.resourcemanager.cosmosdbforpostgresql.models.IdentityType;
import com.azure.resourcemanager.cosmosdbforpostgresql.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Clusters Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-02-preview/examples/
     * ClusterCreateWithAAD.json
     */
    /**
     * Sample code: Create a new cluster with Azure Active Directory Authentication.
     * 
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void createANewClusterWithAzureActiveDirectoryAuthentication(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.clusters().define("testcluster-cmk").withRegion("westus").withExistingResourceGroup("TestGroup")
            .withTags(mapOf())
            .withIdentity(new IdentityProperties().withType(IdentityType.USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-usermanagedidentity",
                    new UserAssignedIdentity())))
            .withAdministratorLoginPassword("password")
            .withDataEncryption(new DataEncryption().withPrimaryKeyUri("fakeTokenPlaceholder")
                .withPrimaryUserAssignedIdentityId(
                    "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-usermanagedidentity")
                .withType(DataEncryptionType.AZURE_KEY_VAULT))
            .withPostgresqlVersion("15").withCitusVersion("12.1").withPreferredPrimaryZone("1")
            .withEnableShardsOnCoordinator(false).withEnableHa(false).withCoordinatorServerEdition("GeneralPurpose")
            .withCoordinatorStorageQuotaInMb(524288).withCoordinatorVCores(4).withCoordinatorEnablePublicIpAccess(true)
            .withNodeServerEdition("MemoryOptimized").withNodeCount(3).withNodeStorageQuotaInMb(524288)
            .withNodeVCores(8).withNodeEnablePublicIpAccess(false).withDatabaseName("citus").create();
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

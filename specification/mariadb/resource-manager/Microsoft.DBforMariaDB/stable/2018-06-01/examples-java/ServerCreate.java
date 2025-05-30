
import com.azure.resourcemanager.mariadb.models.GeoRedundantBackup;
import com.azure.resourcemanager.mariadb.models.MinimalTlsVersionEnum;
import com.azure.resourcemanager.mariadb.models.ServerPropertiesForDefaultCreate;
import com.azure.resourcemanager.mariadb.models.Sku;
import com.azure.resourcemanager.mariadb.models.SkuTier;
import com.azure.resourcemanager.mariadb.models.SslEnforcementEnum;
import com.azure.resourcemanager.mariadb.models.StorageProfile;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Servers Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerCreate.json
     */
    /**
     * Sample code: Create a new server.
     * 
     * @param manager Entry point to MariaDBManager.
     */
    public static void createANewServer(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager.servers().define("mariadbtestsvc4").withRegion("westus").withExistingResourceGroup("testrg")
            .withProperties(new ServerPropertiesForDefaultCreate().withSslEnforcement(SslEnforcementEnum.ENABLED)
                .withMinimalTlsVersion(MinimalTlsVersionEnum.TLS1_2)
                .withStorageProfile(new StorageProfile().withBackupRetentionDays(7)
                    .withGeoRedundantBackup(GeoRedundantBackup.ENABLED).withStorageMB(128000))
                .withAdministratorLogin("cloudsa").withAdministratorLoginPassword("fakeTokenPlaceholder"))
            .withTags(mapOf("ElasticServer", "1"))
            .withSku(
                new Sku().withName("GP_Gen5_2").withTier(SkuTier.GENERAL_PURPOSE).withCapacity(2).withFamily("Gen5"))
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

import com.azure.resourcemanager.postgresql.models.GeoRedundantBackup;
import com.azure.resourcemanager.postgresql.models.MinimalTlsVersionEnum;
import com.azure.resourcemanager.postgresql.models.ServerPropertiesForDefaultCreate;
import com.azure.resourcemanager.postgresql.models.Sku;
import com.azure.resourcemanager.postgresql.models.SkuTier;
import com.azure.resourcemanager.postgresql.models.SslEnforcementEnum;
import com.azure.resourcemanager.postgresql.models.StorageProfile;
import java.util.HashMap;
import java.util.Map;

/** Samples for Servers Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerCreate.json
     */
    /**
     * Sample code: Create a new server.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void createANewServer(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager
            .servers()
            .define("pgtestsvc4")
            .withRegion("westus")
            .withExistingResourceGroup("TestGroup")
            .withProperties(
                new ServerPropertiesForDefaultCreate()
                    .withSslEnforcement(SslEnforcementEnum.ENABLED)
                    .withMinimalTlsVersion(MinimalTlsVersionEnum.TLS1_2)
                    .withStorageProfile(
                        new StorageProfile()
                            .withBackupRetentionDays(7)
                            .withGeoRedundantBackup(GeoRedundantBackup.DISABLED)
                            .withStorageMB(128000))
                    .withAdministratorLogin("cloudsa")
                    .withAdministratorLoginPassword("<administratorLoginPassword>"))
            .withTags(mapOf("ElasticServer", "1"))
            .withSku(new Sku().withName("B_Gen5_2").withTier(SkuTier.BASIC).withCapacity(2).withFamily("Gen5"))
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

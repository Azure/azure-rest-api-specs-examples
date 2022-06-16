import com.azure.resourcemanager.mysql.models.GeoRedundantBackup;
import com.azure.resourcemanager.mysql.models.ServerPropertiesForDefaultCreate;
import com.azure.resourcemanager.mysql.models.Sku;
import com.azure.resourcemanager.mysql.models.SkuTier;
import com.azure.resourcemanager.mysql.models.SslEnforcementEnum;
import com.azure.resourcemanager.mysql.models.StorageProfile;
import java.util.HashMap;
import java.util.Map;

/** Samples for Servers Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerCreate.json
     */
    /**
     * Sample code: Create a new server.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void createANewServer(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .servers()
            .define("mysqltestsvc4")
            .withRegion("westus")
            .withExistingResourceGroup("testrg")
            .withProperties(
                new ServerPropertiesForDefaultCreate()
                    .withSslEnforcement(SslEnforcementEnum.ENABLED)
                    .withStorageProfile(
                        new StorageProfile()
                            .withBackupRetentionDays(7)
                            .withGeoRedundantBackup(GeoRedundantBackup.ENABLED)
                            .withStorageMB(128000))
                    .withAdministratorLogin("cloudsa")
                    .withAdministratorLoginPassword("<administratorLoginPassword>"))
            .withTags(mapOf("ElasticServer", "1"))
            .withSku(
                new Sku().withName("GP_Gen5_2").withTier(SkuTier.GENERAL_PURPOSE).withCapacity(2).withFamily("Gen5"))
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

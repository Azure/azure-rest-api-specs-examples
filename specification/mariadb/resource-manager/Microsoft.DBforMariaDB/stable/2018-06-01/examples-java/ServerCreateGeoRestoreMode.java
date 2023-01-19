import com.azure.resourcemanager.mariadb.models.ServerPropertiesForGeoRestore;
import com.azure.resourcemanager.mariadb.models.Sku;
import com.azure.resourcemanager.mariadb.models.SkuTier;
import java.util.HashMap;
import java.util.Map;

/** Samples for Servers Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerCreateGeoRestoreMode.json
     */
    /**
     * Sample code: Create a server as a geo restore.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void createAServerAsAGeoRestore(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager
            .servers()
            .define("targetserver")
            .withRegion("westus")
            .withExistingResourceGroup("TargetResourceGroup")
            .withProperties(
                new ServerPropertiesForGeoRestore()
                    .withSourceServerId(
                        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/SourceResourceGroup/providers/Microsoft.DBforMariaDB/servers/sourceserver"))
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

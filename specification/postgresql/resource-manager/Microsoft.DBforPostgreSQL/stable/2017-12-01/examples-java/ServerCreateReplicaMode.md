Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-postgresql_1.0.2/sdk/postgresql/azure-resourcemanager-postgresql/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.postgresql.models.ServerPropertiesForReplica;
import com.azure.resourcemanager.postgresql.models.Sku;
import com.azure.resourcemanager.postgresql.models.SkuTier;
import java.util.HashMap;
import java.util.Map;

/** Samples for Servers Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerCreateReplicaMode.json
     */
    /**
     * Sample code: Create a replica server.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void createAReplicaServer(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager
            .servers()
            .define("testserver-replica1")
            .withRegion("westcentralus")
            .withExistingResourceGroup("TestGroup_WestCentralUS")
            .withProperties(
                new ServerPropertiesForReplica()
                    .withSourceServerId(
                        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup_WestCentralUS/providers/Microsoft.DBforPostgreSQL/servers/testserver-master"))
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
```

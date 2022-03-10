Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mysqlflexibleserver_1.0.0-beta.2/sdk/mysqlflexibleserver/azure-resourcemanager-mysqlflexibleserver/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.mysqlflexibleserver.models.Backup;
import com.azure.resourcemanager.mysqlflexibleserver.models.CreateMode;
import com.azure.resourcemanager.mysqlflexibleserver.models.DataEncryption;
import com.azure.resourcemanager.mysqlflexibleserver.models.DataEncryptionType;
import com.azure.resourcemanager.mysqlflexibleserver.models.EnableStatusEnum;
import com.azure.resourcemanager.mysqlflexibleserver.models.HighAvailability;
import com.azure.resourcemanager.mysqlflexibleserver.models.HighAvailabilityMode;
import com.azure.resourcemanager.mysqlflexibleserver.models.Identity;
import com.azure.resourcemanager.mysqlflexibleserver.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.mysqlflexibleserver.models.ServerVersion;
import com.azure.resourcemanager.mysqlflexibleserver.models.Sku;
import com.azure.resourcemanager.mysqlflexibleserver.models.SkuTier;
import com.azure.resourcemanager.mysqlflexibleserver.models.Storage;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

/** Samples for Servers Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ServerCreateWithBYOK.json
     */
    /**
     * Sample code: Create a server with byok.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void createAServerWithByok(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager)
        throws IOException {
        manager
            .servers()
            .define("mysqltestserver")
            .withRegion("southeastasia")
            .withExistingResourceGroup("testrg")
            .withTags(mapOf("num", "1"))
            .withIdentity(
                new Identity()
                    .withType(ManagedServiceIdentityType.USER_ASSIGNED)
                    .withUserAssignedIdentities(
                        mapOf(
                            "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-identity",
                            SerializerFactory
                                .createDefaultManagementSerializerAdapter()
                                .deserialize("{}", Object.class, SerializerEncoding.JSON))))
            .withSku(new Sku().withName("Standard_D2ds_v4").withTier(SkuTier.GENERAL_PURPOSE))
            .withAdministratorLogin("cloudsa")
            .withAdministratorLoginPassword("your_password")
            .withVersion(ServerVersion.FIVE_SEVEN)
            .withAvailabilityZone("1")
            .withCreateMode(CreateMode.DEFAULT)
            .withDataEncryption(
                new DataEncryption()
                    .withPrimaryUserAssignedIdentityId(
                        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-identity")
                    .withPrimaryKeyUri("https://test.vault.azure.net/keys/key/c8a92236622244c0a4fdb892666f671a")
                    .withGeoBackupUserAssignedIdentityId(
                        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-geo-identity")
                    .withGeoBackupKeyUri("https://test-geo.vault.azure.net/keys/key/c8a92236622244c0a4fdb892666f671a")
                    .withType(DataEncryptionType.AZURE_KEY_VAULT))
            .withStorage(new Storage().withStorageSizeGB(100).withIops(600).withAutoGrow(EnableStatusEnum.DISABLED))
            .withBackup(new Backup().withBackupRetentionDays(7).withGeoRedundantBackup(EnableStatusEnum.DISABLED))
            .withHighAvailability(
                new HighAvailability().withMode(HighAvailabilityMode.ZONE_REDUNDANT).withStandbyAvailabilityZone("3"))
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

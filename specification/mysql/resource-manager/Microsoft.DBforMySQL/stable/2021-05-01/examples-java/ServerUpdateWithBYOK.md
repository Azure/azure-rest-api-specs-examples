Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mysqlflexibleserver_1.0.0-beta.2/sdk/mysqlflexibleserver/azure-resourcemanager-mysqlflexibleserver/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.mysqlflexibleserver.models.DataEncryption;
import com.azure.resourcemanager.mysqlflexibleserver.models.DataEncryptionType;
import com.azure.resourcemanager.mysqlflexibleserver.models.Identity;
import com.azure.resourcemanager.mysqlflexibleserver.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.mysqlflexibleserver.models.Server;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

/** Samples for Servers Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ServerUpdateWithBYOK.json
     */
    /**
     * Sample code: Update server with byok.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void updateServerWithByok(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager)
        throws IOException {
        Server resource =
            manager.servers().getByResourceGroupWithResponse("testrg", "mysqltestserver", Context.NONE).getValue();
        resource
            .update()
            .withIdentity(
                new Identity()
                    .withType(ManagedServiceIdentityType.USER_ASSIGNED)
                    .withUserAssignedIdentities(
                        mapOf(
                            "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-identity",
                            SerializerFactory
                                .createDefaultManagementSerializerAdapter()
                                .deserialize("{}", Object.class, SerializerEncoding.JSON))))
            .withDataEncryption(
                new DataEncryption()
                    .withPrimaryUserAssignedIdentityId(
                        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-identity")
                    .withPrimaryKeyUri("https://test.vault.azure.net/keys/key/c8a92236622244c0a4fdb892666f671a")
                    .withGeoBackupUserAssignedIdentityId(
                        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-geo-identity")
                    .withGeoBackupKeyUri("https://test-geo.vault.azure.net/keys/key/c8a92236622244c0a4fdb892666f671a")
                    .withType(DataEncryptionType.AZURE_KEY_VAULT))
            .apply();
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

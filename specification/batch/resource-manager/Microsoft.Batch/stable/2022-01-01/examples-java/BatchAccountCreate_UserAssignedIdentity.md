Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-batch_1.0.0/sdk/batch/azure-resourcemanager-batch/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.batch.models.AutoStorageBaseProperties;
import com.azure.resourcemanager.batch.models.BatchAccountIdentity;
import com.azure.resourcemanager.batch.models.ResourceIdentityType;
import com.azure.resourcemanager.batch.models.UserAssignedIdentities;
import java.util.HashMap;
import java.util.Map;

/** Samples for BatchAccount Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/BatchAccountCreate_UserAssignedIdentity.json
     */
    /**
     * Sample code: BatchAccountCreate_UserAssignedIdentity.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void batchAccountCreateUserAssignedIdentity(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .batchAccounts()
            .define("sampleacct")
            .withRegion("japaneast")
            .withExistingResourceGroup("default-azurebatch-japaneast")
            .withIdentity(
                new BatchAccountIdentity()
                    .withType(ResourceIdentityType.USER_ASSIGNED)
                    .withUserAssignedIdentities(
                        mapOf(
                            "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1",
                            new UserAssignedIdentities())))
            .withAutoStorage(
                new AutoStorageBaseProperties()
                    .withStorageAccountId(
                        "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Storage/storageAccounts/samplestorage"))
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

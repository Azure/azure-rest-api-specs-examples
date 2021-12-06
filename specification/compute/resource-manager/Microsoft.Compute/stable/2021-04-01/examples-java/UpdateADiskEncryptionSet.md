Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.DiskEncryptionSetType;
import com.azure.resourcemanager.compute.models.DiskEncryptionSetUpdate;
import com.azure.resourcemanager.compute.models.KeyForDiskEncryptionSet;
import com.azure.resourcemanager.compute.models.SourceVault;
import java.util.HashMap;
import java.util.Map;

/** Samples for DiskEncryptionSets Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-04-01/examples/UpdateADiskEncryptionSet.json
     */
    /**
     * Sample code: Update a disk encryption set.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateADiskEncryptionSet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDiskEncryptionSets()
            .update(
                "myResourceGroup",
                "myDiskEncryptionSet",
                new DiskEncryptionSetUpdate()
                    .withTags(mapOf("department", "Development", "project", "Encryption"))
                    .withEncryptionType(DiskEncryptionSetType.ENCRYPTION_AT_REST_WITH_CUSTOMER_KEY)
                    .withActiveKey(
                        new KeyForDiskEncryptionSet()
                            .withSourceVault(
                                new SourceVault()
                                    .withId(
                                        "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"))
                            .withKeyUrl("https://myvmvault.vault-int.azure-int.net/keys/keyName/keyVersion")),
                Context.NONE);
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

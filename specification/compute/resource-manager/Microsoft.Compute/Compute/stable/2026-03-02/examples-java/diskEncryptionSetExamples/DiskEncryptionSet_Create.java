
import com.azure.resourcemanager.compute.fluent.models.DiskEncryptionSetInner;
import com.azure.resourcemanager.compute.models.DiskEncryptionSetIdentityType;
import com.azure.resourcemanager.compute.models.DiskEncryptionSetType;
import com.azure.resourcemanager.compute.models.EncryptionSetIdentity;
import com.azure.resourcemanager.compute.models.KeyForDiskEncryptionSet;
import com.azure.resourcemanager.compute.models.SourceVault;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DiskEncryptionSets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskEncryptionSetExamples/DiskEncryptionSet_Create.json
     */
    /**
     * Sample code: create a disk encryption set.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void createADiskEncryptionSet(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDiskEncryptionSets().createOrUpdate("myResourceGroup", "myDiskEncryptionSet",
            new DiskEncryptionSetInner().withLocation("West US")
                .withIdentity(new EncryptionSetIdentity().withType(DiskEncryptionSetIdentityType.SYSTEM_ASSIGNED))
                .withEncryptionType(DiskEncryptionSetType.ENCRYPTION_AT_REST_WITH_CUSTOMER_KEY)
                .withActiveKey(new KeyForDiskEncryptionSet().withSourceVault(new SourceVault().withId(
                    "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"))
                    .withKeyUrl("fakeTokenPlaceholder")),
            com.azure.core.util.Context.NONE);
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

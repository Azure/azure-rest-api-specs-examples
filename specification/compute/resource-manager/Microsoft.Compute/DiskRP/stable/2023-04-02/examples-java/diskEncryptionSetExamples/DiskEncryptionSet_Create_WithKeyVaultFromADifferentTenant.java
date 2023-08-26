import com.azure.resourcemanager.compute.fluent.models.DiskEncryptionSetInner;
import com.azure.resourcemanager.compute.models.DiskEncryptionSetIdentityType;
import com.azure.resourcemanager.compute.models.DiskEncryptionSetType;
import com.azure.resourcemanager.compute.models.EncryptionSetIdentity;
import com.azure.resourcemanager.compute.models.KeyForDiskEncryptionSet;
import com.azure.resourcemanager.compute.models.VirtualMachineIdentityUserAssignedIdentities;
import java.util.HashMap;
import java.util.Map;

/** Samples for DiskEncryptionSets CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-04-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_Create_WithKeyVaultFromADifferentTenant.json
     */
    /**
     * Sample code: Create a disk encryption set with key vault from a different tenant.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createADiskEncryptionSetWithKeyVaultFromADifferentTenant(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDiskEncryptionSets()
            .createOrUpdate(
                "myResourceGroup",
                "myDiskEncryptionSet",
                new DiskEncryptionSetInner()
                    .withLocation("West US")
                    .withIdentity(
                        new EncryptionSetIdentity()
                            .withType(DiskEncryptionSetIdentityType.USER_ASSIGNED)
                            .withUserAssignedIdentities(
                                mapOf(
                                    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}",
                                    new VirtualMachineIdentityUserAssignedIdentities())))
                    .withEncryptionType(DiskEncryptionSetType.ENCRYPTION_AT_REST_WITH_CUSTOMER_KEY)
                    .withActiveKey(new KeyForDiskEncryptionSet().withKeyUrl("fakeTokenPlaceholder"))
                    .withFederatedClientId("00000000-0000-0000-0000-000000000000"),
                com.azure.core.util.Context.NONE);
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

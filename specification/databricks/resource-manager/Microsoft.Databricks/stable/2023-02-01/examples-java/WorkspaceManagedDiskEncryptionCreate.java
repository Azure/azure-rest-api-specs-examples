import com.azure.resourcemanager.databricks.models.EncryptionEntitiesDefinition;
import com.azure.resourcemanager.databricks.models.EncryptionKeySource;
import com.azure.resourcemanager.databricks.models.ManagedDiskEncryption;
import com.azure.resourcemanager.databricks.models.ManagedDiskEncryptionKeyVaultProperties;
import com.azure.resourcemanager.databricks.models.WorkspacePropertiesEncryption;
import java.util.HashMap;
import java.util.Map;

/** Samples for Workspaces CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/WorkspaceManagedDiskEncryptionCreate.json
     */
    /**
     * Sample code: Create a workspace with Customer-Managed Key (CMK) encryption for Managed Disks.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void createAWorkspaceWithCustomerManagedKeyCMKEncryptionForManagedDisks(
        com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager
            .workspaces()
            .define("myWorkspace")
            .withRegion("westus")
            .withExistingResourceGroup("rg")
            .withManagedResourceGroupId("/subscriptions/subid/resourceGroups/myManagedRG")
            .withEncryption(
                new WorkspacePropertiesEncryption()
                    .withEntities(
                        new EncryptionEntitiesDefinition()
                            .withManagedDisk(
                                new ManagedDiskEncryption()
                                    .withKeySource(EncryptionKeySource.MICROSOFT_KEYVAULT)
                                    .withKeyVaultProperties(
                                        new ManagedDiskEncryptionKeyVaultProperties()
                                            .withKeyVaultUri("fakeTokenPlaceholder")
                                            .withKeyName("fakeTokenPlaceholder")
                                            .withKeyVersion("fakeTokenPlaceholder"))
                                    .withRotationToLatestKeyVersionEnabled(true))))
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

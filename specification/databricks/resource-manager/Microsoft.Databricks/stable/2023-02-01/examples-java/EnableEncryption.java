import com.azure.resourcemanager.databricks.models.Encryption;
import com.azure.resourcemanager.databricks.models.KeySource;
import com.azure.resourcemanager.databricks.models.WorkspaceCustomBooleanParameter;
import com.azure.resourcemanager.databricks.models.WorkspaceCustomParameters;
import com.azure.resourcemanager.databricks.models.WorkspaceEncryptionParameter;
import java.util.HashMap;
import java.util.Map;

/** Samples for Workspaces CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/EnableEncryption.json
     */
    /**
     * Sample code: Enable Customer-Managed Key (CMK) encryption on a workspace which is prepared for encryption.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void enableCustomerManagedKeyCMKEncryptionOnAWorkspaceWhichIsPreparedForEncryption(
        com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager
            .workspaces()
            .define("myWorkspace")
            .withRegion("westus")
            .withExistingResourceGroup("rg")
            .withManagedResourceGroupId("/subscriptions/subid/resourceGroups/myManagedRG")
            .withParameters(
                new WorkspaceCustomParameters()
                    .withPrepareEncryption(new WorkspaceCustomBooleanParameter().withValue(true))
                    .withEncryption(
                        new WorkspaceEncryptionParameter()
                            .withValue(
                                new Encryption()
                                    .withKeySource(KeySource.MICROSOFT_KEYVAULT)
                                    .withKeyName("fakeTokenPlaceholder")
                                    .withKeyVersion("fakeTokenPlaceholder")
                                    .withKeyVaultUri("fakeTokenPlaceholder"))))
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

import com.azure.resourcemanager.databricks.models.Encryption;
import com.azure.resourcemanager.databricks.models.KeySource;
import com.azure.resourcemanager.databricks.models.WorkspaceCustomBooleanParameter;
import com.azure.resourcemanager.databricks.models.WorkspaceCustomParameters;
import com.azure.resourcemanager.databricks.models.WorkspaceEncryptionParameter;

/** Samples for Workspaces CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/EnableEncryption.json
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
}

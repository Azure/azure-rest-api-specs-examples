import com.azure.resourcemanager.databricks.models.Encryption;
import com.azure.resourcemanager.databricks.models.KeySource;
import com.azure.resourcemanager.databricks.models.WorkspaceCustomParameters;
import com.azure.resourcemanager.databricks.models.WorkspaceEncryptionParameter;

/** Samples for Workspaces CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/DisableEncryption.json
     */
    /**
     * Sample code: Revert Customer-Managed Key (CMK) encryption to Microsoft Managed Keys encryption on a workspace.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void revertCustomerManagedKeyCMKEncryptionToMicrosoftManagedKeysEncryptionOnAWorkspace(
        com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager
            .workspaces()
            .define("myWorkspace")
            .withRegion("westus")
            .withExistingResourceGroup("rg")
            .withManagedResourceGroupId("/subscriptions/subid/resourceGroups/myManagedRG")
            .withParameters(
                new WorkspaceCustomParameters()
                    .withEncryption(
                        new WorkspaceEncryptionParameter()
                            .withValue(new Encryption().withKeySource(KeySource.DEFAULT))))
            .create();
    }
}

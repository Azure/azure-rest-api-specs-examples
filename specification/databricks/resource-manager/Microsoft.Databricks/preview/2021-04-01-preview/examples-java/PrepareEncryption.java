import com.azure.resourcemanager.databricks.models.WorkspaceCustomBooleanParameter;
import com.azure.resourcemanager.databricks.models.WorkspaceCustomParameters;

/** Samples for Workspaces CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/PrepareEncryption.json
     */
    /**
     * Sample code: Create a workspace which is ready for Customer-Managed Key (CMK) encryption.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void createAWorkspaceWhichIsReadyForCustomerManagedKeyCMKEncryption(
        com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager
            .workspaces()
            .define("myWorkspace")
            .withRegion("westus")
            .withExistingResourceGroup("rg")
            .withManagedResourceGroupId("/subscriptions/subid/resourceGroups/myManagedRG")
            .withParameters(
                new WorkspaceCustomParameters()
                    .withPrepareEncryption(new WorkspaceCustomBooleanParameter().withValue(true)))
            .create();
    }
}

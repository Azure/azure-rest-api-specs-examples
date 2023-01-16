import com.azure.resourcemanager.databricks.models.WorkspaceCustomParameters;
import com.azure.resourcemanager.databricks.models.WorkspaceCustomStringParameter;

/** Samples for Workspaces CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/WorkspaceCreateWithParameters.json
     */
    /**
     * Sample code: Create or update workspace with custom parameters.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void createOrUpdateWorkspaceWithCustomParameters(
        com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager
            .workspaces()
            .define("myWorkspace")
            .withRegion("westus")
            .withExistingResourceGroup("rg")
            .withManagedResourceGroupId("/subscriptions/subid/resourceGroups/myManagedRG")
            .withParameters(
                new WorkspaceCustomParameters()
                    .withCustomVirtualNetworkId(
                        new WorkspaceCustomStringParameter()
                            .withValue(
                                "/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/myNetwork"))
                    .withCustomPublicSubnetName(new WorkspaceCustomStringParameter().withValue("myPublicSubnet"))
                    .withCustomPrivateSubnetName(new WorkspaceCustomStringParameter().withValue("myPrivateSubnet")))
            .create();
    }
}

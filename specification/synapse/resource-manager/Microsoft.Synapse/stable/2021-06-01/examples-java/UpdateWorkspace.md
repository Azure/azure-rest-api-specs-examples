```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.CustomerManagedKeyDetails;
import com.azure.resourcemanager.synapse.models.EncryptionDetails;
import com.azure.resourcemanager.synapse.models.ManagedIdentity;
import com.azure.resourcemanager.synapse.models.ManagedVirtualNetworkSettings;
import com.azure.resourcemanager.synapse.models.PurviewConfiguration;
import com.azure.resourcemanager.synapse.models.ResourceIdentityType;
import com.azure.resourcemanager.synapse.models.Workspace;
import com.azure.resourcemanager.synapse.models.WorkspaceKeyDetails;
import com.azure.resourcemanager.synapse.models.WorkspacePublicNetworkAccess;
import com.azure.resourcemanager.synapse.models.WorkspaceRepositoryConfiguration;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Workspaces Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/UpdateWorkspace.json
     */
    /**
     * Sample code: Update a workspace.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void updateAWorkspace(com.azure.resourcemanager.synapse.SynapseManager manager) {
        Workspace resource =
            manager
                .workspaces()
                .getByResourceGroupWithResponse("resourceGroup1", "workspace1", Context.NONE)
                .getValue();
        resource
            .update()
            .withTags(mapOf("key", "value"))
            .withIdentity(new ManagedIdentity().withType(ResourceIdentityType.SYSTEM_ASSIGNED))
            .withSqlAdministratorLoginPassword("password")
            .withManagedVirtualNetworkSettings(
                new ManagedVirtualNetworkSettings()
                    .withPreventDataExfiltration(false)
                    .withLinkedAccessCheckOnTargetResource(false)
                    .withAllowedAadTenantIdsForLinking(Arrays.asList("740239CE-A25B-485B-86A0-262F29F6EBDB")))
            .withWorkspaceRepositoryConfiguration(
                new WorkspaceRepositoryConfiguration()
                    .withType("FactoryGitHubConfiguration")
                    .withHostname("")
                    .withAccountName("adifferentacount")
                    .withProjectName("myproject")
                    .withRepositoryName("myrepository")
                    .withCollaborationBranch("master")
                    .withRootFolder("/"))
            .withPurviewConfiguration(
                new PurviewConfiguration()
                    .withPurviewResourceId(
                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.ProjectPurview/accounts/accountname1"))
            .withEncryption(
                new EncryptionDetails()
                    .withCmk(
                        new CustomerManagedKeyDetails()
                            .withKey(
                                new WorkspaceKeyDetails()
                                    .withName("default")
                                    .withKeyVaultUrl("https://vault.azure.net/keys/key1"))))
            .withPublicNetworkAccess(WorkspacePublicNetworkAccess.ENABLED)
            .apply();
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

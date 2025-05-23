
import com.azure.resourcemanager.synapse.models.CspWorkspaceAdminProperties;
import com.azure.resourcemanager.synapse.models.CustomerManagedKeyDetails;
import com.azure.resourcemanager.synapse.models.DataLakeStorageAccountDetails;
import com.azure.resourcemanager.synapse.models.EncryptionDetails;
import com.azure.resourcemanager.synapse.models.KekIdentityProperties;
import com.azure.resourcemanager.synapse.models.ManagedIdentity;
import com.azure.resourcemanager.synapse.models.ManagedVirtualNetworkSettings;
import com.azure.resourcemanager.synapse.models.PurviewConfiguration;
import com.azure.resourcemanager.synapse.models.ResourceIdentityType;
import com.azure.resourcemanager.synapse.models.UserAssignedManagedIdentity;
import com.azure.resourcemanager.synapse.models.WorkspaceKeyDetails;
import com.azure.resourcemanager.synapse.models.WorkspacePublicNetworkAccess;
import com.azure.resourcemanager.synapse.models.WorkspaceRepositoryConfiguration;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Workspaces CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateWorkspace.json
     */
    /**
     * Sample code: Create or update a workspace.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void createOrUpdateAWorkspace(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.workspaces().define("workspace1").withRegion("East US").withExistingResourceGroup("resourceGroup1")
            .withTags(mapOf("key", "fakeTokenPlaceholder"))
            .withIdentity(new ManagedIdentity().withType(ResourceIdentityType.SYSTEM_ASSIGNED_USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1",
                    new UserAssignedManagedIdentity())))
            .withDefaultDataLakeStorage(new DataLakeStorageAccountDetails()
                .withAccountUrl("https://accountname.dfs.core.windows.net").withFilesystem("default"))
            .withSqlAdministratorLoginPassword("password")
            .withManagedResourceGroupName("workspaceManagedResourceGroupUnique").withSqlAdministratorLogin("login")
            .withManagedVirtualNetwork("default")
            .withEncryption(new EncryptionDetails().withCmk(new CustomerManagedKeyDetails()
                .withKey(new WorkspaceKeyDetails().withName("default").withKeyVaultUrl("fakeTokenPlaceholder"))
                .withKekIdentity(new KekIdentityProperties().withUserAssignedIdentity(
                    "/subscriptions/b64d7b94-73e7-4d36-94b2-7764ea3fd74a/resourcegroups/SynapseCI/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1")
                    .withUseSystemAssignedIdentity(false))))
            .withManagedVirtualNetworkSettings(new ManagedVirtualNetworkSettings().withPreventDataExfiltration(false)
                .withLinkedAccessCheckOnTargetResource(false)
                .withAllowedAadTenantIdsForLinking(Arrays.asList("740239CE-A25B-485B-86A0-262F29F6EBDB")))
            .withWorkspaceRepositoryConfiguration(
                new WorkspaceRepositoryConfiguration().withType("FactoryGitHubConfiguration").withHostname("")
                    .withAccountName("mygithubaccount").withProjectName("myproject").withRepositoryName("myrepository")
                    .withCollaborationBranch("master").withRootFolder("/"))
            .withPurviewConfiguration(new PurviewConfiguration().withPurviewResourceId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.ProjectPurview/accounts/accountname1"))
            .withPublicNetworkAccess(WorkspacePublicNetworkAccess.ENABLED)
            .withCspWorkspaceAdminProperties(new CspWorkspaceAdminProperties()
                .withInitialWorkspaceAdminObjectId("6c20646f-8050-49ec-b3b1-80a0e58e454d"))
            .create();
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

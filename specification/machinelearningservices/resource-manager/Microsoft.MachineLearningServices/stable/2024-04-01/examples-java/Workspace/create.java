
import com.azure.resourcemanager.machinelearning.models.EncryptionKeyVaultProperties;
import com.azure.resourcemanager.machinelearning.models.EncryptionProperty;
import com.azure.resourcemanager.machinelearning.models.EncryptionStatus;
import com.azure.resourcemanager.machinelearning.models.IdentityForCmk;
import com.azure.resourcemanager.machinelearning.models.ManagedServiceIdentity;
import com.azure.resourcemanager.machinelearning.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.machinelearning.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.machinelearning.models.SharedPrivateLinkResource;
import com.azure.resourcemanager.machinelearning.models.UserAssignedIdentity;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Workspaces CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/create.json
     */
    /**
     * Sample code: Create Workspace.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createWorkspace(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.workspaces().define("testworkspace").withExistingResourceGroup("workspace-1234")
            .withRegion("eastus2euap")
            .withIdentity(new ManagedServiceIdentity()
                .withType(ManagedServiceIdentityType.SYSTEM_ASSIGNED_USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testuai",
                    new UserAssignedIdentity())))
            .withDescription("test description").withFriendlyName("HelloName")
            .withKeyVault(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.KeyVault/vaults/testkv")
            .withApplicationInsights(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/microsoft.insights/components/testinsights")
            .withContainerRegistry(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ContainerRegistry/registries/testRegistry")
            .withStorageAccount(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/accountcrud-1234/providers/Microsoft.Storage/storageAccounts/testStorageAccount")
            .withEncryption(new EncryptionProperty().withStatus(EncryptionStatus.ENABLED)
                .withIdentity(new IdentityForCmk().withUserAssignedIdentity(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testuai"))
                .withKeyVaultProperties(new EncryptionKeyVaultProperties().withKeyVaultArmId("fakeTokenPlaceholder")
                    .withKeyIdentifier("fakeTokenPlaceholder").withIdentityClientId("")))
            .withHbiWorkspace(false)
            .withSharedPrivateLinkResources(Arrays.asList(new SharedPrivateLinkResource().withName("testdbresource")
                .withPrivateLinkResourceId(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.DocumentDB/databaseAccounts/testdbresource/privateLinkResources/Sql")
                .withGroupId("Sql").withRequestMessage("Please approve")
                .withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)))
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

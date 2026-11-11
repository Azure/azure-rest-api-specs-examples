
import com.azure.resourcemanager.chaos.models.Workspace;
import com.azure.resourcemanager.chaos.models.WorkspaceCustomerManagedKeyEncryptionPatchMergePatchUpdateOrCreate;
import com.azure.resourcemanager.chaos.models.WorkspaceEncryptionPatchMergePatchUpdateOrCreate;
import com.azure.resourcemanager.chaos.models.WorkspaceUpdateProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Workspaces Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/Workspaces_Update_ReplaceCustomerKey.json
     */
    /**
     * Sample code: Update only the customer key URI of an encrypted Workspace.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void
        updateOnlyTheCustomerKeyURIOfAnEncryptedWorkspace(com.azure.resourcemanager.chaos.ChaosManager manager) {
        Workspace resource = manager.workspaces()
            .getByResourceGroupWithResponse("exampleRG", "exampleWorkspace", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update()
            .withProperties(new WorkspaceUpdateProperties()
                .withEncryption(new WorkspaceEncryptionPatchMergePatchUpdateOrCreate().withCustomerManagedKeyEncryption(
                    new WorkspaceCustomerManagedKeyEncryptionPatchMergePatchUpdateOrCreate()
                        .withKeyEncryptionKeyUrl("fakeTokenPlaceholder"))))
            .apply();
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

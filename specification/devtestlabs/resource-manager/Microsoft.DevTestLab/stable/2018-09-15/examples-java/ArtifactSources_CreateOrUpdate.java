
import com.azure.resourcemanager.devtestlabs.models.EnableStatus;
import com.azure.resourcemanager.devtestlabs.models.SourceControlType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ArtifactSources CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/
     * ArtifactSources_CreateOrUpdate.json
     */
    /**
     * Sample code: ArtifactSources_CreateOrUpdate.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void artifactSourcesCreateOrUpdate(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.artifactSources().define("{artifactSourceName}").withRegion((String) null)
            .withExistingLab("resourceGroupName", "{labName}").withTags(mapOf("tagName1", "tagValue1"))
            .withDisplayName("{displayName}").withUri("{artifactSourceUri}")
            .withSourceType(SourceControlType.fromString("{VsoGit|GitHub|StorageAccount}"))
            .withFolderPath("{folderPath}").withArmTemplateFolderPath("{armTemplateFolderPath}")
            .withBranchRef("{branchRef}").withSecurityToken("{securityToken}")
            .withStatus(EnableStatus.fromString("{Enabled|Disabled}")).create();
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

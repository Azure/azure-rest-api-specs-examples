
import com.azure.resourcemanager.securityinsights.models.MetadataAuthor;
import com.azure.resourcemanager.securityinsights.models.MetadataModel;

/**
 * Samples for Metadata Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/metadata/PatchMetadata.json
     */
    /**
     * Sample code: Update metadata.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void updateMetadata(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        MetadataModel resource = manager.metadatas()
            .getWithResponse("myRg", "myWorkspace", "metadataName", com.azure.core.util.Context.NONE).getValue();
        resource.update().withAuthor(new MetadataAuthor().withName("User Name").withEmail("email@microsoft.com"))
            .apply();
    }
}

import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.MetadataAuthor;
import com.azure.resourcemanager.securityinsights.models.MetadataModel;

/** Samples for Metadata Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/metadata/PatchMetadata.json
     */
    /**
     * Sample code: Update metadata.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void updateMetadata(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        MetadataModel resource =
            manager.metadatas().getWithResponse("myRg", "myWorkspace", "metadataName", Context.NONE).getValue();
        resource
            .update()
            .withAuthor(new MetadataAuthor().withName("User Name").withEmail("email@microsoft.com"))
            .apply();
    }
}

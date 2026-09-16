
import com.azure.resourcemanager.securityinsights.models.RepositoryAccessKind;
import com.azure.resourcemanager.securityinsights.models.RepositoryAccessProperties;

/**
 * Samples for SourceControls Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/sourcecontrols/DeleteSourceControl.json
     */
    /**
     * Sample code: Delete a source control.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        deleteASourceControl(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.sourceControls().deleteWithResponse("myRg", "myWorkspace", "789e0c1f-4a3d-43ad-809c-e713b677b04a",
            new RepositoryAccessProperties().withKind(RepositoryAccessKind.OAUTH).withCode("fakeTokenPlaceholder")
                .withState("state").withClientId("54b3c2c0-1f48-4a1c-af9f-6399c3240b73"),
            com.azure.core.util.Context.NONE);
    }
}

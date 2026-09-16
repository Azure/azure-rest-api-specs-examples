
/**
 * Samples for HuntComments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/hunts/DeleteHuntComment.json
     */
    /**
     * Sample code: Delete a hunt comment.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteAHuntComment(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.huntComments().deleteWithResponse("myRg", "myWorkspace", "163e7b2a-a2ec-4041-aaba-d878a38f265f",
            "2216d0e1-91e3-4902-89fd-d2df8c123456", com.azure.core.util.Context.NONE);
    }
}

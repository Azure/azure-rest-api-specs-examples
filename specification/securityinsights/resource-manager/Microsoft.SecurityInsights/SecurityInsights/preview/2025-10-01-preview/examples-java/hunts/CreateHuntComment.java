
/**
 * Samples for HuntComments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/hunts/CreateHuntComment.json
     */
    /**
     * Sample code: Creates or updates a hunt comment.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        createsOrUpdatesAHuntComment(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.huntComments().define("2216d0e1-91e3-4902-89fd-d2df8c535096")
            .withExistingHunt("myRg", "myWorkspace", "163e7b2a-a2ec-4041-aaba-d878a38f265f")
            .withMessage("This is a test comment.").create();
    }
}

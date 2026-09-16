
/**
 * Samples for HuntComments Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/hunts/GetHuntCommentById.json
     */
    /**
     * Sample code: Get a hunt comment.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAHuntComment(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.huntComments().getWithResponse("myRg", "myWorkspace", "163e7b2a-a2ec-4041-aaba-d878a38f265f",
            "2216d0e1-91e3-4902-89fd-d2df8c535096", com.azure.core.util.Context.NONE);
    }
}

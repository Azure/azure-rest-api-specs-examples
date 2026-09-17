
/**
 * Samples for HuntRelations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/hunts/DeleteHuntRelation.json
     */
    /**
     * Sample code: Delete a hunt relation.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteAHuntRelation(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.huntRelations().deleteWithResponse("myRg", "myWorkspace", "163e7b2a-a2ec-4041-aaba-d878a38f265f",
            "2216d0e1-91e3-4902-89fd-d2df8c535096", com.azure.core.util.Context.NONE);
    }
}

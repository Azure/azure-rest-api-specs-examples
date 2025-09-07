
/**
 * Samples for SpotPlacementScores Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-05/GetSpotPlacementScores.json
     */
    /**
     * Sample code: Gets the metadata of Spot Placement Scores.
     * 
     * @param manager Entry point to ComputeRecommenderManager.
     */
    public static void getsTheMetadataOfSpotPlacementScores(
        com.azure.resourcemanager.computerecommender.ComputeRecommenderManager manager) {
        manager.spotPlacementScores().getWithResponse("eastus", com.azure.core.util.Context.NONE);
    }
}


import com.azure.resourcemanager.mongocluster.models.CheckNameAvailabilityRequest;

/**
 * Samples for MongoClusters CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01-preview/MongoClusters_NameAvailability_AlreadyExists.json
     */
    /**
     * Sample code: Checks and returns that the Mongo Cluster name is already in-use.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void checksAndReturnsThatTheMongoClusterNameIsAlreadyInUse(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.mongoClusters()
            .checkNameAvailabilityWithResponse("westus2", new CheckNameAvailabilityRequest()
                .withName("existingmongocluster").withType("Microsoft.DocumentDB/mongoClusters"),
                com.azure.core.util.Context.NONE);
    }
}

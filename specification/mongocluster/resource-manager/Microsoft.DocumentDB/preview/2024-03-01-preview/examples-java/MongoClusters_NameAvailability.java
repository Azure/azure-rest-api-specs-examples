
import com.azure.resourcemanager.mongocluster.models.CheckNameAvailabilityRequest;

/**
 * Samples for MongoClusters CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-01-preview/MongoClusters_NameAvailability.json
     */
    /**
     * Sample code: Checks and confirms the Mongo Cluster name is availability for use.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void checksAndConfirmsTheMongoClusterNameIsAvailabilityForUse(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.mongoClusters().checkNameAvailabilityWithResponse("westus2", new CheckNameAvailabilityRequest()
            .withName("newmongocluster").withType("Microsoft.DocumentDB/mongoClusters"),
            com.azure.core.util.Context.NONE);
    }
}

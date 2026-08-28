
import com.azure.resourcemanager.horizondb.models.HorizonDbReplicaProperties;
import com.azure.resourcemanager.horizondb.models.ReplicaRole;

/**
 * Samples for HorizonDbReplicas CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Replicas_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a HorizonDB replica.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void createOrUpdateAHorizonDBReplica(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbReplicas().define("examplereplica")
            .withExistingPool("exampleresourcegroup", "examplecluster", "examplepool")
            .withProperties(new HorizonDbReplicaProperties().withRole(ReplicaRole.READ).withAvailabilityZone("1"))
            .create();
    }
}

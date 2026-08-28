
import com.azure.resourcemanager.horizondb.models.HorizonDbReplica;
import com.azure.resourcemanager.horizondb.models.HorizonDbReplicaPropertiesForPatchUpdate;
import com.azure.resourcemanager.horizondb.models.ReplicaRole;

/**
 * Samples for HorizonDbReplicas Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Replicas_Update.json
     */
    /**
     * Sample code: Update a HorizonDB replica.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void updateAHorizonDBReplica(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        HorizonDbReplica resource = manager.horizonDbReplicas().getWithResponse("exampleresourcegroup",
            "examplecluster", "examplepool", "examplereplica", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withProperties(new HorizonDbReplicaPropertiesForPatchUpdate().withRole(ReplicaRole.READ_WRITE)).apply();
    }
}

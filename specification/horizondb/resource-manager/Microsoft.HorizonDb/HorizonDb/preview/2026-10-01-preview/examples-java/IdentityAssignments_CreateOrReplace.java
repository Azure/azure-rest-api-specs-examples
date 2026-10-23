
import com.azure.resourcemanager.horizondb.models.ChangeDataCaptureIdentityAssignmentProperties;

/**
 * Samples for IdentityAssignments CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/IdentityAssignments_CreateOrReplace.json
     */
    /**
     * Sample code: Create a change data capture identity assignment for a UAMI attached to the parent cluster.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void createAChangeDataCaptureIdentityAssignmentForAUAMIAttachedToTheParentCluster(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.identityAssignments().define("fabric-writer")
            .withExistingChangeDataCaptureDestination("example-resource-group", "example-cluster", "fabric-mirroring")
            .withProperties(new ChangeDataCaptureIdentityAssignmentProperties().withIdentityResourceId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/identity-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/fabric-writer"))
            .create();
    }
}

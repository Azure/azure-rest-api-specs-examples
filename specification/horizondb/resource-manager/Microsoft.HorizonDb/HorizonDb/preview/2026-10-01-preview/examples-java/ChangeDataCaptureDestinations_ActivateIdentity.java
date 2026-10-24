
import com.azure.resourcemanager.horizondb.models.ActivateChangeDataCaptureIdentityRequest;

/**
 * Samples for ChangeDataCaptureDestinations ActivateIdentity.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/ChangeDataCaptureDestinations_ActivateIdentity.json
     */
    /**
     * Sample code: Activate a validated identity assignment for a change data capture destination.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void activateAValidatedIdentityAssignmentForAChangeDataCaptureDestination(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.changeDataCaptureDestinations().activateIdentity("example-resource-group", "*", "example-cluster",
            "fabric-mirroring",
            new ActivateChangeDataCaptureIdentityRequest().withIdentityAssignmentResourceId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/changeDataCaptureDestinations/fabric-mirroring/identityAssignments/fabric-writer")
                .withDrainExistingSessions(true),
            com.azure.core.util.Context.NONE);
    }
}

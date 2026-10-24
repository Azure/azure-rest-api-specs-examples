
/**
 * Samples for IdentityAssignments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/IdentityAssignments_Delete.json
     */
    /**
     * Sample code: Delete an inactive change data capture identity assignment.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void deleteAnInactiveChangeDataCaptureIdentityAssignment(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.identityAssignments().delete("example-resource-group", "example-cluster", "fabric-mirroring",
            "fabric-writer", com.azure.core.util.Context.NONE);
    }
}

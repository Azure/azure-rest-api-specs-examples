
/**
 * Samples for IdentityAssignments List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/IdentityAssignments_List.json
     */
    /**
     * Sample code: List change data capture identity assignments.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        listChangeDataCaptureIdentityAssignments(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.identityAssignments().list("example-resource-group", "example-cluster", "fabric-mirroring",
            com.azure.core.util.Context.NONE);
    }
}

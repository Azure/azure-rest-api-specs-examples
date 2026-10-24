
/**
 * Samples for IdentityAssignments Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/IdentityAssignments_Get.json
     */
    /**
     * Sample code: Get a change data capture identity assignment.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        getAChangeDataCaptureIdentityAssignment(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.identityAssignments().getWithResponse("example-resource-group", "example-cluster", "fabric-mirroring",
            "fabric-writer", com.azure.core.util.Context.NONE);
    }
}

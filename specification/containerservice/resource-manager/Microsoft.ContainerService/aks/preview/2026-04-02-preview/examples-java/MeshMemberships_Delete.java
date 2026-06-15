
/**
 * Samples for MeshMemberships Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/MeshMemberships_Delete.json
     */
    /**
     * Sample code: Delete Mesh Membership.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        deleteMeshMembership(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getMeshMemberships().delete("rg1", "clustername1", "meshmembership1",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for MeshMemberships Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/MeshMemberships_Get.json
     */
    /**
     * Sample code: Get Mesh Membership.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void getMeshMembership(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getMeshMemberships().getWithResponse("rg1", "clustername1", "meshmembership1",
            com.azure.core.util.Context.NONE);
    }
}

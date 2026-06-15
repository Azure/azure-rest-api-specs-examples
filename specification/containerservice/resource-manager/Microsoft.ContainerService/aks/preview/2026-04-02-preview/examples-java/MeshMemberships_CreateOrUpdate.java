
import com.azure.resourcemanager.containerservice.fluent.models.MeshMembershipInner;
import com.azure.resourcemanager.containerservice.models.MeshMembershipProperties;

/**
 * Samples for MeshMemberships CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/MeshMemberships_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update Mesh Membership.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        createOrUpdateMeshMembership(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getMeshMemberships().createOrUpdate("rg1", "clustername1", "meshmembership1",
            new MeshMembershipInner().withProperties(new MeshMembershipProperties().withManagedMeshId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.AppLink/applinks/applink1/appLinkMembers/member1")),
            com.azure.core.util.Context.NONE);
    }
}

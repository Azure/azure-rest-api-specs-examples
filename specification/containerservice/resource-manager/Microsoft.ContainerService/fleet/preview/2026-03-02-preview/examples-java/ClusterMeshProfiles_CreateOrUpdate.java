
import com.azure.resourcemanager.containerservicefleet.models.MemberSelector;

/**
 * Samples for ClusterMeshProfiles CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/ClusterMeshProfiles_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a ClusterMeshProfile.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void createOrUpdateAClusterMeshProfile(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.clusterMeshProfiles().define("clustermeshprofile1").withExistingFleet("rgfleets", "fleet1")
            .withMemberSelector(new MemberSelector().withByLabel("env=production")).withIfMatch("uktvayathbu").create();
    }
}

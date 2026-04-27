
import com.azure.resourcemanager.containerservice.models.ManagedClusterAadProfile;

/**
 * Samples for ManagedClusters ResetAADProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ManagedClustersResetAADProfile.json
     */
    /**
     * Sample code: Reset AAD Profile.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void resetAADProfile(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().resetAADProfile("rg1", "clustername1",
            new ManagedClusterAadProfile().withClientAppId("clientappid").withServerAppId("serverappid")
                .withServerAppSecret("fakeTokenPlaceholder").withTenantId("tenantid"),
            com.azure.core.util.Context.NONE);
    }
}

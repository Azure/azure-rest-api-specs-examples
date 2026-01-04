
import com.azure.resourcemanager.containerservice.models.ManagedClusterAadProfile;

/**
 * Samples for ManagedClusters ResetAadProfile.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * ManagedClustersResetAADProfile.json
     */
    /**
     * Sample code: Reset AAD Profile.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void resetAADProfile(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getManagedClusters().resetAadProfile("rg1", "clustername1",
            new ManagedClusterAadProfile().withClientAppId("clientappid").withServerAppId("serverappid")
                .withServerAppSecret("fakeTokenPlaceholder").withTenantId("tenantid"),
            com.azure.core.util.Context.NONE);
    }
}

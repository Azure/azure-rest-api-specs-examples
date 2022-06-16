import com.azure.core.util.Context;
import com.azure.resourcemanager.containerservice.models.ManagedClusterAadProfile;

/** Samples for ManagedClusters ResetAadProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-02-01/examples/ManagedClustersResetAADProfile.json
     */
    /**
     * Sample code: Reset AAD Profile.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void resetAADProfile(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getManagedClusters()
            .resetAadProfile(
                "rg1",
                "clustername1",
                new ManagedClusterAadProfile()
                    .withClientAppId("clientappid")
                    .withServerAppId("serverappid")
                    .withServerAppSecret("serverappsecret")
                    .withTenantId("tenantid"),
                Context.NONE);
    }
}

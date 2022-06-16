import com.azure.core.util.Context;
import com.azure.resourcemanager.containerservice.models.ManagedClusterServicePrincipalProfile;

/** Samples for ManagedClusters ResetServicePrincipalProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-10-01/examples/ManagedClustersResetServicePrincipalProfile.json
     */
    /**
     * Sample code: Reset Service Principal Profile.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void resetServicePrincipalProfile(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getManagedClusters()
            .resetServicePrincipalProfile(
                "rg1",
                "clustername1",
                new ManagedClusterServicePrincipalProfile().withClientId("clientid").withSecret("secret"),
                Context.NONE);
    }
}

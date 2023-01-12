import com.azure.resourcemanager.hybridkubernetes.models.AuthenticationMethod;
import com.azure.resourcemanager.hybridkubernetes.models.ListClusterUserCredentialProperties;

/** Samples for ConnectedCluster ListClusterUserCredential. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/stable/2021-10-01/examples/ConnectedClustersListClusterCredentialResultHPToken.json
     */
    /**
     * Sample code: ListClusterUserCredentialNonAadCSPExample.
     *
     * @param manager Entry point to HybridKubernetesManager.
     */
    public static void listClusterUserCredentialNonAadCSPExample(
        com.azure.resourcemanager.hybridkubernetes.HybridKubernetesManager manager) {
        manager
            .connectedClusters()
            .listClusterUserCredentialWithResponse(
                "k8sc-rg",
                "testCluster",
                new ListClusterUserCredentialProperties()
                    .withAuthenticationMethod(AuthenticationMethod.TOKEN)
                    .withClientProxy(false),
                com.azure.core.util.Context.NONE);
    }
}

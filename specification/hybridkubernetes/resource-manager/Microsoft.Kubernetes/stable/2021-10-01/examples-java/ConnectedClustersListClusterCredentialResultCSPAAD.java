import com.azure.resourcemanager.hybridkubernetes.models.AuthenticationMethod;
import com.azure.resourcemanager.hybridkubernetes.models.ListClusterUserCredentialProperties;

/** Samples for ConnectedCluster ListClusterUserCredential. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/stable/2021-10-01/examples/ConnectedClustersListClusterCredentialResultCSPAAD.json
     */
    /**
     * Sample code: ListClusterUserCredentialExample.
     *
     * @param manager Entry point to HybridKubernetesManager.
     */
    public static void listClusterUserCredentialExample(
        com.azure.resourcemanager.hybridkubernetes.HybridKubernetesManager manager) {
        manager
            .connectedClusters()
            .listClusterUserCredentialWithResponse(
                "k8sc-rg",
                "testCluster",
                new ListClusterUserCredentialProperties()
                    .withAuthenticationMethod(AuthenticationMethod.AAD)
                    .withClientProxy(true),
                com.azure.core.util.Context.NONE);
    }
}

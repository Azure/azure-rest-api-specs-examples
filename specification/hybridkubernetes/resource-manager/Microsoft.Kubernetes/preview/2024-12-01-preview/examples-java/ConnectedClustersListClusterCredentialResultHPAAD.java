
import com.azure.resourcemanager.hybridkubernetes.models.AuthenticationMethod;
import com.azure.resourcemanager.hybridkubernetes.models.ListClusterUserCredentialProperties;

/**
 * Samples for ConnectedCluster ListClusterUserCredential.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/preview/2024-12-01-preview/examples/
     * ConnectedClustersListClusterCredentialResultHPAAD.json
     */
    /**
     * Sample code: ListClusterUserCredentialCSPExample.
     * 
     * @param manager Entry point to HybridKubernetesManager.
     */
    public static void listClusterUserCredentialCSPExample(
        com.azure.resourcemanager.hybridkubernetes.HybridKubernetesManager manager) {
        manager.connectedClusters().listClusterUserCredentialWithResponse(
            "k8sc-rg", "testCluster", new ListClusterUserCredentialProperties()
                .withAuthenticationMethod(AuthenticationMethod.AAD).withClientProxy(false),
            com.azure.core.util.Context.NONE);
    }
}

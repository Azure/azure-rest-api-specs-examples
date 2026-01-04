
import com.azure.resourcemanager.containerservice.fluent.models.PrivateLinkResourceInner;

/**
 * Samples for ResolvePrivateLinkServiceId Post.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * ResolvePrivateLinkServiceId.json
     */
    /**
     * Sample code: Resolve the Private Link Service ID for Managed Cluster.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        resolveThePrivateLinkServiceIDForManagedCluster(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getResolvePrivateLinkServiceIds().postWithResponse("rg1",
            "clustername1", new PrivateLinkResourceInner().withName("management"), com.azure.core.util.Context.NONE);
    }
}

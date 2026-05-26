
import com.azure.resourcemanager.containerservice.fluent.models.PrivateLinkResourceInner;

/**
 * Samples for ResolvePrivateLinkServiceId POST.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/ResolvePrivateLinkServiceId.json
     */
    /**
     * Sample code: Resolve the Private Link Service ID for Managed Cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void resolveThePrivateLinkServiceIDForManagedCluster(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getResolvePrivateLinkServiceIds().pOSTWithResponse("rg1", "clustername1",
            new PrivateLinkResourceInner().withName("management"), com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for JWTAuthenticators ListByManagedCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/JWTAuthenticators_List.json
     */
    /**
     * Sample code: List JWT authenticators by Managed Cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void listJWTAuthenticatorsByManagedCluster(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getJWTAuthenticators().listByManagedCluster("rg1", "clustername1",
            com.azure.core.util.Context.NONE);
    }
}

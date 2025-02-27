
/**
 * Samples for PublicCloudConnectors TestPermissions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/PublicCloudConnectors_TestPermissions.json
     */
    /**
     * Sample code: PublicCloudConnectors_TestPermissions.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void publicCloudConnectorsTestPermissions(
        com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.publicCloudConnectors().testPermissions("rgpublicCloud", "rzygvnpsnrdylwzdbsscjazvamyxmh",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for PublicCloudConnectors ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/PublicCloudConnectors_ListByResourceGroup.json
     */
    /**
     * Sample code: PublicCloudConnectors_ListByResourceGroup.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void publicCloudConnectorsListByResourceGroup(
        com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.publicCloudConnectors().listByResourceGroup("rgpublicCloud", com.azure.core.util.Context.NONE);
    }
}

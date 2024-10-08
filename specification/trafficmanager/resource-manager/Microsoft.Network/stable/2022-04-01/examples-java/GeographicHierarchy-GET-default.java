
/**
 * Samples for GeographicHierarchies GetDefault.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/GeographicHierarchy-
     * GET-default.json
     */
    /**
     * Sample code: GeographicHierarchy-GET-default.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void geographicHierarchyGETDefault(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.trafficManagerProfiles().manager().serviceClient().getGeographicHierarchies()
            .getDefaultWithResponse(com.azure.core.util.Context.NONE);
    }
}

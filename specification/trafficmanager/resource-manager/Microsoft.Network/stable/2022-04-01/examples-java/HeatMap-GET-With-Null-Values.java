/** Samples for HeatMap Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/HeatMap-GET-With-Null-Values.json
     */
    /**
     * Sample code: HeatMap-GET-With-Null-Values.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void heatMapGETWithNullValues(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .trafficManagerProfiles()
            .manager()
            .serviceClient()
            .getHeatMaps()
            .getWithResponse(
                "azuresdkfornetautoresttrafficmanager1323",
                "azuresdkfornetautoresttrafficmanager3880",
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}

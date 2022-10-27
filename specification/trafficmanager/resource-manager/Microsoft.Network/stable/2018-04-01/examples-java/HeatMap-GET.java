import com.azure.core.util.Context;

/** Samples for HeatMap Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-04-01/examples/HeatMap-GET.json
     */
    /**
     * Sample code: HeatMap-GET.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void heatMapGET(com.azure.resourcemanager.AzureResourceManager azure) {
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
                Context.NONE);
    }
}

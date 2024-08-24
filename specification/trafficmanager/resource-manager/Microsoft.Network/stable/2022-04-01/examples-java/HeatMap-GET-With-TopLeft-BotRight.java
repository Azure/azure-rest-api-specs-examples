
import java.util.Arrays;

/**
 * Samples for HeatMap Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/HeatMap-GET-With-
     * TopLeft-BotRight.json
     */
    /**
     * Sample code: HeatMap-GET-With-TopLeft-BotRight.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void heatMapGETWithTopLeftBotRight(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.trafficManagerProfiles().manager().serviceClient().getHeatMaps().getWithResponse(
            "azuresdkfornetautoresttrafficmanager1323", "azuresdkfornetautoresttrafficmanager3880",
            Arrays.asList(10.0, 50.001), Arrays.asList(-50.001, 80.0), com.azure.core.util.Context.NONE);
    }
}

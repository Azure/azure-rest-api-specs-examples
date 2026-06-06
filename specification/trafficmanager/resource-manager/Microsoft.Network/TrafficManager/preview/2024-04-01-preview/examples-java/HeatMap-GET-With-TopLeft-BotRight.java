
import java.util.Arrays;

/**
 * Samples for HeatMap Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-04-01-preview/HeatMap-GET-With-TopLeft-BotRight.json
     */
    /**
     * Sample code: HeatMap-GET-With-TopLeft-BotRight.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void heatMapGETWithTopLeftBotRight(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getHeatMaps().getWithResponse("azuresdkfornetautoresttrafficmanager1323",
            "azuresdkfornetautoresttrafficmanager3880", Arrays.asList(10.0, 50.001), Arrays.asList(-50.001, 80.0),
            com.azure.core.util.Context.NONE);
    }
}

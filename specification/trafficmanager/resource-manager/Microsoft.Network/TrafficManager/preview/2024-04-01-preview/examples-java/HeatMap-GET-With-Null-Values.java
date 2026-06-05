
/**
 * Samples for HeatMap Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-04-01-preview/HeatMap-GET-With-Null-Values.json
     */
    /**
     * Sample code: HeatMap-GET-With-Null-Values.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void heatMapGETWithNullValues(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getHeatMaps().getWithResponse("azuresdkfornetautoresttrafficmanager1323",
            "azuresdkfornetautoresttrafficmanager3880", null, null, com.azure.core.util.Context.NONE);
    }
}

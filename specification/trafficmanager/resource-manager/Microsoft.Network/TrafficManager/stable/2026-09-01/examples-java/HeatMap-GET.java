
/**
 * Samples for HeatMap Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/HeatMap-GET.json
     */
    /**
     * Sample code: HeatMap-GET.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void heatMapGET(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getHeatMaps().getWithResponse("azuresdkfornetautoresttrafficmanager1323",
            "azuresdkfornetautoresttrafficmanager3880", null, null, com.azure.core.util.Context.NONE);
    }
}

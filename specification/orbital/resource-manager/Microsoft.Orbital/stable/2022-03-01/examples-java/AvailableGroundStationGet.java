import com.azure.core.util.Context;

/** Samples for AvailableGroundStations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/AvailableGroundStationGet.json
     */
    /**
     * Sample code: Get GroundStation.
     *
     * @param manager Entry point to OrbitalManager.
     */
    public static void getGroundStation(com.azure.resourcemanager.orbital.OrbitalManager manager) {
        manager.availableGroundStations().getWithResponse("EASTUS2_0", Context.NONE);
    }
}

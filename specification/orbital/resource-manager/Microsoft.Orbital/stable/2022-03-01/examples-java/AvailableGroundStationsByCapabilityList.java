import com.azure.core.util.Context;
import com.azure.resourcemanager.orbital.models.CapabilityParameter;

/** Samples for AvailableGroundStations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/AvailableGroundStationsByCapabilityList.json
     */
    /**
     * Sample code: List of Ground Stations by Capability.
     *
     * @param manager Entry point to OrbitalManager.
     */
    public static void listOfGroundStationsByCapability(com.azure.resourcemanager.orbital.OrbitalManager manager) {
        manager.availableGroundStations().list(CapabilityParameter.EARTH_OBSERVATION, Context.NONE);
    }
}

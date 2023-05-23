/** Samples for FrontDoors List. */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorListAll.json
     */
    /**
     * Sample code: List all Front Doors.
     *
     * @param manager Entry point to FrontDoorManager.
     */
    public static void listAllFrontDoors(com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager.frontDoors().list(com.azure.core.util.Context.NONE);
    }
}

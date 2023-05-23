/** Samples for FrontDoors Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorDelete.json
     */
    /**
     * Sample code: Delete Front Door.
     *
     * @param manager Entry point to FrontDoorManager.
     */
    public static void deleteFrontDoor(com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager.frontDoors().delete("rg1", "frontDoor1", com.azure.core.util.Context.NONE);
    }
}

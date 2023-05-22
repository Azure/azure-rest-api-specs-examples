/** Samples for FrontDoors ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorList.json
     */
    /**
     * Sample code: List Front Doors in a Resource Group.
     *
     * @param manager Entry point to FrontDoorManager.
     */
    public static void listFrontDoorsInAResourceGroup(com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager.frontDoors().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}

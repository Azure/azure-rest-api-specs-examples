
/**
 * Samples for Capacities ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2021-01-01/listCapacitiesInResourceGroup.json
     */
    /**
     * Sample code: List capacities in resource group.
     * 
     * @param manager Entry point to PowerBIDedicatedManager.
     */
    public static void
        listCapacitiesInResourceGroup(com.azure.resourcemanager.powerbidedicated.PowerBIDedicatedManager manager) {
        manager.capacities().listByResourceGroup("TestRG", com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for ServiceFabricSchedules List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/ServiceFabricSchedules_List.json
     */
    /**
     * Sample code: ServiceFabricSchedules_List.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void serviceFabricSchedulesList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.serviceFabricSchedules().list("resourceGroupName", "{labName}", "@me", "{serviceFrabicName}", null,
            null, null, null, com.azure.core.util.Context.NONE);
    }
}

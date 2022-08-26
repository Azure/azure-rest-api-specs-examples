import com.azure.core.util.Context;

/** Samples for Schedules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Schedules/getSchedule.json
     */
    /**
     * Sample code: getSchedule.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void getSchedule(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.schedules().getWithResponse("testrg123", "testlab", "schedule1", Context.NONE);
    }
}

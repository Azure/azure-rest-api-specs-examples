
/**
 * Samples for JobStream Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/job/getJobStream.json
     */
    /**
     * Sample code: Get job stream.
     * 
     * @param manager Entry point to AutomationManager.
     */
    public static void getJobStream(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.jobStreams().getWithResponse("mygroup", "ContoseAutomationAccount", "foo",
            "851b2101-686f-40e2-8a4b-5b8df08afbd1_00636535684910693884_00000000000000000001", null,
            com.azure.core.util.Context.NONE);
    }
}

import com.azure.core.util.Context;

/** Samples for TestJobStreams Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/getTestJobStream.json
     */
    /**
     * Sample code: Get test job stream.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getTestJobStream(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .testJobStreams()
            .getWithResponse(
                "mygroup",
                "ContoseAutomationAccount",
                "Get-AzureVMTutorial",
                "851b2101-686f-40e2-8a4b-5b8df08afbd1_00636535684910693884_00000000000000000001",
                Context.NONE);
    }
}

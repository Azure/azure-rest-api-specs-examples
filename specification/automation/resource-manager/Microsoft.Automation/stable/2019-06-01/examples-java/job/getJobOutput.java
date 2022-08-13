import com.azure.core.util.Context;

/** Samples for Job GetOutput. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/job/getJobOutput.json
     */
    /**
     * Sample code: Get Job Output.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getJobOutput(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.jobs().getOutputWithResponse("mygroup", "ContoseAutomationAccount", "foo", null, Context.NONE);
    }
}

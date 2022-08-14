import com.azure.core.util.Context;

/** Samples for SourceControl Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/sourceControl/getSourceControl.json
     */
    /**
     * Sample code: Get a source control.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getASourceControl(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.sourceControls().getWithResponse("rg", "sampleAccount9", "sampleSourceControl", Context.NONE);
    }
}

import com.azure.core.util.Context;

/** Samples for SourceControl Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/sourceControl/deleteSourceControl.json
     */
    /**
     * Sample code: Delete a source control.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void deleteASourceControl(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.sourceControls().deleteWithResponse("rg", "sampleAccount9", "sampleSourceControl", Context.NONE);
    }
}

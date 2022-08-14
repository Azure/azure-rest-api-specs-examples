import com.azure.core.util.Context;

/** Samples for Variable Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/deleteVariable.json
     */
    /**
     * Sample code: Delete a variable.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void deleteAVariable(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.variables().deleteWithResponse("rg", "sampleAccount9", "sampleVariable", Context.NONE);
    }
}

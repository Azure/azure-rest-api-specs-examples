import com.azure.core.util.Context;

/** Samples for Variable Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getVariable.json
     */
    /**
     * Sample code: Get a variable.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getAVariable(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.variables().getWithResponse("rg", "sampleAccount9", "sampleVariable", Context.NONE);
    }
}

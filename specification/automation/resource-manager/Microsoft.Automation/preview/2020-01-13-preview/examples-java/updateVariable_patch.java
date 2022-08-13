import com.azure.core.util.Context;
import com.azure.resourcemanager.automation.models.Variable;

/** Samples for Variable Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/updateVariable_patch.json
     */
    /**
     * Sample code: Update a variable.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void updateAVariable(com.azure.resourcemanager.automation.AutomationManager manager) {
        Variable resource =
            manager.variables().getWithResponse("rg", "sampleAccount9", "sampleVariable", Context.NONE).getValue();
        resource.update().withName("sampleVariable").withValue("\"ComputerName3.domain.com\"").apply();
    }
}

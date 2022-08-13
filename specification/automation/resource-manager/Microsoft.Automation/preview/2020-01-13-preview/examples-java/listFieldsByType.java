import com.azure.core.util.Context;

/** Samples for ObjectDataTypes ListFieldsByType. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listFieldsByType.json
     */
    /**
     * Sample code: Get a list of fields of a given type across all accessible modules.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getAListOfFieldsOfAGivenTypeAcrossAllAccessibleModules(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.objectDataTypes().listFieldsByType("rg", "MyAutomationAccount", "MyCustomType", Context.NONE);
    }
}

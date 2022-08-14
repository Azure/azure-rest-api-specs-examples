import com.azure.core.util.Context;

/** Samples for Fields ListByType. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listFieldsByModuleAndType.json
     */
    /**
     * Sample code: Get a list of fields of a given type.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getAListOfFieldsOfAGivenType(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.fields().listByType("rg", "MyAutomationAccount", "MyModule", "MyCustomType", Context.NONE);
    }
}

import com.azure.core.util.Context;

/** Samples for Python2Package Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/deletePython2Package.json
     */
    /**
     * Sample code: Delete a python 2 package.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void deleteAPython2Package(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .python2Packages()
            .deleteWithResponse("rg", "myAutomationAccount33", "OmsCompositeResources", Context.NONE);
    }
}

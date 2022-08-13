import com.azure.core.util.Context;

/** Samples for Python2Package Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getPython2Package.json
     */
    /**
     * Sample code: Get a python 2 package.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getAPython2Package(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.python2Packages().getWithResponse("rg", "myAutomationAccount33", "OmsCompositeResources", Context.NONE);
    }
}

import com.azure.core.util.Context;
import com.azure.resourcemanager.automation.fluent.models.GraphicalRunbookContentInner;

/** Samples for ResourceProvider ConvertGraphRunbookContent. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/serializeGraphRunbookContent.json
     */
    /**
     * Sample code: Get Graphical raw runbook content from graphical runbook JSON object.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getGraphicalRawRunbookContentFromGraphicalRunbookJSONObject(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .resourceProviders()
            .convertGraphRunbookContentWithResponse(
                "rg",
                "MyAutomationAccount",
                new GraphicalRunbookContentInner().withGraphRunbookJson("<GraphRunbookJSON>"),
                Context.NONE);
    }
}
